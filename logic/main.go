package main

import (
	"fmt"
	"gmimo/common/log"
	"gmimo/logic/config"
	"gmimo/logic/controller"
	"gmimo/logic/model"
	"gmimo/logic/server"
	"os"
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

	if err = model.InitRedisClient(config.Redis); err != nil {
		log.Error("Redis server 连接失败")
		return
	}
	log.Info("Redis server 连接成功")

	if err = controller.InitKafkaProducer(config.Kafka); err != nil {
		log.Error("Kafka Producer创建失败")
		return
	}
	log.Info("Kafka Producer创建成功")

	return
}

func main() {
	if err := Init(); err != nil {
		fmt.Println("项目初始化失败:", err)
		os.Exit(-1)
	}

	// 启动RPC server
	go server.RunRpcServer()

	// 启动http server
	httpServer := server.NewHttpServer()

	// 监听系统中断信号
	server.WatchShutdown(httpServer)
}
