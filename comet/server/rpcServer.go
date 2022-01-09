package server

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"gmimo/comet/config"
	"gmimo/comet/stub"
	"gmimo/common/log"
	"gmimo/common/middleware"
	messageProto "gmimo/common/proto/message"
	"google.golang.org/grpc"
	"net"
	"os"
)

func RunRpcServer() {
	// 创建server
	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain( // 加载Logger，Recovery中间件
			middleware.RpcLogger,
			middleware.RpcRecovery),
	}
	srv := grpc.NewServer(opts...)

	// 注册服务
	messageProto.RegisterSendMessageServer(srv, &stub.MessageStub{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Rpc.CometPort))
	if err != nil {
		log.Error("rpc server is closed:" + err.Error())
		os.Exit(-1)
	}

	log.Info("rpc server is running")
	if err := srv.Serve(lis); err != nil {
		log.Error("rpc server is closed:", err.Error())
		os.Exit(-1)
	}
}
