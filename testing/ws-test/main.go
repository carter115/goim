package main

import (
	"flag"
	"fmt"
	"gmimo/common/log"
	"os"
	"os/signal"
	"sync"
	"time"
)

var (
	// 配置参数
	wsAddr string // websocket连接地址
	conns  int    // 连接数
	preUid string // 用户名前缀
	roomId string //房间ID

	forever   bool // 是否循环发送房间消息
	interval  int  // 下一步骤间隔时间
	totalTime int  // 持续发送消息多少秒后停止程序

	waitTime int // 每条消息的间隔等待时间
	waitN    int // 添加时间误差范围

	content    string // 消息体
	keeplive   bool   // 保持心跳
	onlyWsconn bool   // 只保持连接
	level      string //打印日志级别

	startTime int64 // 用于统计TPS
	endTime   int64
)

// ----------------------初始化----------------------
func initFlag() {
	//flag.StringVar(&wsAddr, "addr", "ws://192.168.30.88:18080/ws/connect", "websocket url ")
	//flag.StringVar(&wsAddr, "addr", "ws://47.74.183.125:18080/ws/connect", "websocket url ")
	//flag.StringVar(&wsAddr, "addr", "wss://mimotest.elelive.cn:18081/ws/connect", "websocket连接地址")

	flag.StringVar(&wsAddr, "addr", "wss://mimo-ws.elelive.net/ws/connect", "websocket连接地址")
	//flag.StringVar(&wsAddr, "addr", "ws://localhost:10000/ws/connect", "websocket连接地址")
	flag.IntVar(&conns, "conns", 2, "连接用户数")
	flag.StringVar(&preUid, "uid", "UserA-", "用户名前缀")
	flag.StringVar(&roomId, "room", "room001", "房间ID")

	flag.BoolVar(&forever, "forever", true, "是否循环发送房间消息.")
	flag.IntVar(&interval, "interval", 5, "每个步骤的间隔时间,需要等待连接、加房间等.单位:秒")
	flag.IntVar(&totalTime, "time", 0, "持续发送消息N秒后,停止程序")
	flag.IntVar(&waitTime, "wait", 5000, "发送消息的间隔时间.单位:毫秒")
	flag.IntVar(&waitN, "waitn", 2000, "添加时间误差范围,减少并发的压力.单位:毫秒")

	flag.StringVar(&content, "content", "", "消息内容")
	flag.BoolVar(&keeplive, "keeplive", true, "是否发送心跳包")
	flag.BoolVar(&onlyWsconn, "onlyWsconn", false, "只保持websocket连接")

	flag.StringVar(&level, "level", "info", "输出日志级别")
	flag.Parse()
	fmt.Printf("命令行参数: \naddr:[%s], conns:[%d], uid:[%s], room:[%s], forever:[%t], interval:[%d], time:[%d], wait:[%d], waitn:[%d], content:[%s], keeplive:[%t], onlyWsconn:[%t], level:[%s]\n",
		wsAddr, conns, preUid, roomId, forever, interval, totalTime, waitTime, waitN, content, keeplive, onlyWsconn, level)
}

func initLog() {
	conf := log.LogConfig{
		Level:       level,
		StashServer: "192.168.0.240:4560",
		Hooks:       []string{"stash"},
		Outputs:     []string{"stdout"},
	}
	if err := log.InitLogger(conf, "mimo-test"); err != nil {
		fmt.Println("初始化日志模块失败:", err)
		os.Exit(-1)
	}
}

func main() {
	// 加载命令行参数
	initFlag()

	// 加载日志组件
	initLog()

	var (
		wg        sync.WaitGroup
		userConns = make([]*WsConnectionClient, conns)
		quit      = make(chan os.Signal, 1)
	)

	// panic recovery
	defer func() {
		if e := recover(); e != nil {
			log.Error("panic:", e)
		}
	}()

	go watchShutdown(userConns, quit) // 处理程序的中断信号

	// 创建N个连接
	wg.Add(conns)
	fmt.Println("-> 正在创建用户连接")
	for i := 0; i < conns; i++ {
		go NewWsConnectionClient(&wg, userConns, i)
	}
	wg.Wait()
	fmt.Printf("-> 创建%d个用户连接\n", len(userConns))

	// keeplive:只发心跳，onlyWsconn:只保持连接
	if keeplive {
		for _, ws := range userConns {
			if ws != nil {
				go ws.Keeplive()
			}
		}
		// 只做连接
		if onlyWsconn {
			fmt.Println("-> 正在保持连接...")
			select {} // 死循环
		}
	}

	//// 加入房间
	//for _, ws := range userConns {
	//	if ws != nil {
	//		go ws.JoinRoom()
	//	}
	//}
	//fmt.Printf("-> 正在加入房间,等待%d秒\n", interval)
	time.Sleep(time.Duration(interval) * time.Second)

	// 1. 循环发送消息
	fmt.Println("-> 正在发送房间消息:", time.Now())
	startTime = time.Now().Unix() // 开始计时
	if forever {
		// 持续发送多长时间后,停止程序
		if totalTime > 0 {
			ti := time.Duration(totalTime) * time.Second
			fmt.Printf("-> 注意: 程序在 %v 后停止运行\n", ti)
			time.AfterFunc(ti, func() {
				quit <- os.Interrupt
			})
		}
		startTime = time.Now().Unix() // 重置计时

		// 1.1 多个连接循环发消息
		if len(content) > 0 {
			for _, ws := range userConns {
				if ws != nil {
					go func(w *WsConnectionClient) {
						w.JoinRoom() // 加房间
						// 创建房间消息
						msg := NewRoomMessage(roomId, content)
						w.WriteLoop(msg)
					}(ws)
				}
			}
		}

		select {} // 死循环,代码不往下走

	} else {
		// 1.2 只发送一条
		for _, ws := range userConns {
			if ws != nil {
				ws.JoinRoom() // 加房间
				msg := NewRoomMessage(roomId, content)
				go ws.writeMessage(msg)
			}
		}
		time.Sleep(time.Duration(interval) * time.Second) // 等待发消息

		//  结束程序，统计收到的消息
		quit <- os.Interrupt
		fmt.Println("-> 结束发送房间消息", time.Now())
	}
}

// ----------------------功能----------------------

// 关闭连接，并统计结果
func statClosed(WsConns []*WsConnectionClient) {
	var total int64
	runSecond := endTime - startTime
	if runSecond < 1 { // 设置最小1秒
		runSecond = 1
	}

	for _, ws := range WsConns {
		go ws.Close() // 关闭连接
		total += ws.RecvCount
		log.Infof("[%s] 收到 %d 条消息, TPS:%.2f, Time:%v",
			ws.Uid, ws.RecvCount, float32(ws.RecvCount)/float32(runSecond), runSecond)
	}
	log.Infof("所有消息统计: 数量:%d, TPS:%d, 时间:%v", total, total/runSecond, runSecond)
	fmt.Printf("0. 程序正在停止.请等待%d秒\n", interval)
	time.Sleep(time.Duration(interval) * time.Second) // 等待异步写日志
}

// 监控进程的中断信号
func watchShutdown(WsConns []*WsConnectionClient, quit chan os.Signal) {
	signal.Notify(quit, os.Interrupt)
	<-quit
	endTime = time.Now().Unix() // 结束计时
	log.Error("server is closed.")
	statClosed(WsConns) // 统计结果
	os.Exit(-1)
}
