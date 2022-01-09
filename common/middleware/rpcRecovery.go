package middleware

import (
	"context"
	"fmt"
	"gmimo/common/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RpcRecovery(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if err := recover(); err != nil {
			// 打印错误日志
			log.Errorc(ctx, fmt.Sprintf("PANIC: %v", err))
			// TODO panic 错误，统一包装成json数据，进行返回
			err = status.Errorf(codes.Internal, "Panic err: %v", err)
		}
	}()

	return handler(ctx, req)
}
