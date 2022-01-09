package main

import (
	"fmt"
	"gmimo/common/log"
	"gmimo/push/config"
	"gmimo/push/controller"
	"gmimo/push/model"
	"os"
	"os/signal"
	"time"
)

// 项目初始化
func Init() (err error) {
	if err = config.InitConfig(); err != nil {
		fmt.Println("初始化配置模块失败:", err)
		return
	}
	fmt.Println("成功加载配置模块", config.String())

	if err = log.InitLogger(config.Logger, config.App.Name); err != nil {
		fmt.Println("初始化日志模块失败:", err)
		return
	}
	fmt.Println("成功加载日志器")

	controller.InitScheduler() // 1. 启动消息调度器
	if controller.Scheduler == nil {
		log.Error("message scheduler创建失败")
		return
	}

	if err = model.InitRedisClient(config.Redis); err != nil {
		log.Error("Redis启动失败")
		return
	}
	log.Info("Redis连接成功")

	if err = model.InitMongoClient(config.Mongo); err != nil {
		log.Error("Mongo启动失败")
		return
	}
	log.Info("Mongo连接成功")

	if err = controller.InitConsumerGroup(config.Kafka); err != nil {
		log.Error("Kafka ConsumerGroup创建失败")
		return
	}
	log.Info("Kafka ConsumerGroup创建成功")
	go controller.KafkaPull() // 2. 消费消息

	return
}

func main() {
	if err := Init(); err != nil {
		fmt.Println("项目初始化失败:", err)
		os.Exit(-1)
	}

	// 程序退出
	var quit = make(chan os.Signal)
	signal.Notify(quit, os.Interrupt) //监听中断信号
	<-quit                            //收到中断信号
	log.Error("server is shutdown")
	time.Sleep(time.Second)
}
