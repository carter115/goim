package main

import (
	"fmt"
	"gmimo/comet/config"
	"gmimo/comet/server"
	"gmimo/common/log"
)

// @title Gmimo Server
// @version 0.1
// @description IM服务
func main() {
	if err := config.InitConfig(); err != nil {
		fmt.Println("初始化配置模块失败:", err)
		return
	}
	fmt.Println("成功加载配置模块", config.String())

	if err := log.InitLogger(config.Logger, config.App.Name); err != nil {
		fmt.Println("初始化日志模块失败:", err)
		return
	}
	fmt.Println("成功加载日志器")

	// 启动RPC server
	go server.RunRpcServer()

	// 启动http server
	httpServer := server.NewHttpServer()

	// 监听系统中断信号
	server.WatchShutdown(httpServer)
}
