package server

import (
	"context"
	"gmimo/comet/config"
	"gmimo/comet/router"
	"gmimo/common/log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func NewHttpServer() *http.Server {
	// 创建Gin Engine
	engine := router.InitRouter()

	// 创建Http Server
	server := &http.Server{
		Addr:           config.App.Addr,
		ReadTimeout:    config.App.ReadTimeout,
		WriteTimeout:   config.App.WriteTimeout,
		MaxHeaderBytes: config.App.MaxKB,
		Handler:        engine,
	}

	// 启动http server
	log.Info("http server is running")
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Errorf("http server is closed: %v", err)
			os.Exit(-1)
		}
	}()
	return server
}

func WatchShutdown(server *http.Server) {
	var quit = make(chan os.Signal)
	signal.Notify(quit, os.Interrupt) //监听中断信号
	<-quit                            //收到中断信号

	// 停止http server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Warn("shutdown server error:", err)
	}

	time.Sleep(time.Second) // 休眠1秒，等待异步写日志
	log.Error("server is shutdown")
}
