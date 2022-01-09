package connection

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"gmimo/comet/stub"
	"gmimo/common/e"
	"gmimo/common/log"
	roomProto "gmimo/common/proto/room"
	"sync"
	"time"
)

// 管理用户的websocket连接池
var WsConnPoolMap sync.Map

// 自定义心跳的间隔时间
var interval = 5 * time.Second

// 客户端连接
type WsConnection struct {
	Uid    string
	Mid    string
	Token  string
	WsConn *websocket.Conn

	inChan    chan []byte // 读队列
	outChan   chan []byte // 写队列
	mutex     sync.Mutex  // 避免重复关闭管道
	isClosed  bool
	closeChan chan byte // 关闭通知

}

func InitWsConnection(uid, mid string, wsConn *websocket.Conn) (conn *WsConnection, err error) {
	conn = &WsConnection{
		Uid:       uid,
		Mid:       mid,
		WsConn:    wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}

	// 启动读协程
	go conn.readLoop()
	// 启动写协程
	go conn.writeLoop()

	// 自定义的心跳
	//go conn.heartbeat()
	return

}

// 内部实现
func (conn *WsConnection) readLoop() {
	var (
		data []byte
		err  error
	)
	defer e.Recovery()
	for {
		if _, data, err = conn.WsConn.ReadMessage(); err != nil {
			goto ERR
		}

		//阻塞在这里，等待inChan有空闲位置
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			goto ERR
		}
	}
ERR:
	conn.Close()
}

func (conn *WsConnection) writeLoop() {
	var (
		data []byte
		err  error
	)
	defer e.Recovery()
	for {
		select {
		case data = <-conn.outChan:
			//log.Info("comet write message:", data)
		case <-conn.closeChan:
			goto ERR
		}
		if err = conn.WsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}

func (conn *WsConnection) Close() {
	// 线程安全，可多次调用
	conn.WsConn.Close()

	// 利用标记，让closeChan只关闭一次
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true

		// 移除pool中的用户连接
		if _, ok := WsConnPoolMap.Load(conn.Uid); ok {
			WsConnPoolMap.Delete(conn.Uid)
			log.Infof("[%s] connection is closed", conn.Uid)
		}

		// 发布一条离开房间的订阅消息
		event := UserConnEvent{
			Uid:    conn.Uid,
			ConnId: fmt.Sprintf("%p", conn),
			Action: LeaveRoom,
		}
		go Publish(event)

		// 离开房间
		req := &roomProto.Request{Mid: conn.Mid, Uid: conn.Uid}
		stub.CallRoomLeave(context.Background(), req)
	}
	conn.mutex.Unlock()
}

// 启动线程，不断发消息
func (conn *WsConnection) heartbeat() {
	var err error
	for {
		if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
			fmt.Println("heartbeat error:", err)
			return
		}
		time.Sleep(interval)
	}
}

// 以下是API
func (conn *WsConnection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("ws connection is closeed")
	}
	return
}

func (conn *WsConnection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("ws connection is closeed")
	}
	return
}

// 查询连接池的用户
func ConnPoolMapList() (m map[string]interface{}) {
	m = make(map[string]interface{})
	WsConnPoolMap.Range(func(key, value interface{}) bool {
		k := fmt.Sprintf("%v", key)
		v := fmt.Sprintf("%p", value)
		m[k] = v
		return true
	})
	return
}

// 获取WsConnection
func LoadUserWsConnection(uid string) *WsConnection {
	val, ok := WsConnPoolMap.Load(uid)
	if ok {
		conn := val.(*WsConnection)
		return conn
	}
	return nil
}

// 获取某个房间的WsConnection
func LoadRoomWsConnection(mid string) (conns []*WsConnection) {
	WsConnPoolMap.Range(func(key, value interface{}) bool {
		v := value.(*WsConnection)
		if mid == v.Mid {
			conns = append(conns, v)
		}
		return true
	})
	return
}

// TODO 获取房间连接的uid
