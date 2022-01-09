package server

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"gmimo/common/log"
	"gmimo/common/middleware"
	helloProto "gmimo/common/proto/hello"
	messageProto "gmimo/common/proto/message"
	roomProto "gmimo/common/proto/room"
	"gmimo/logic/config"
	"gmimo/logic/stub"
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
	helloProto.RegisterSayHelloServer(srv, &stub.HelloStub{})
	messageProto.RegisterSendMessageServer(srv, &stub.MessageStub{})
	roomProto.RegisterRoomServer(srv, &stub.RoomStub{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Rpc.LogicPort))
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
