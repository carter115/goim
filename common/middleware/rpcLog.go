package middleware

import (
	"context"
	"fmt"
	"gmimo/common/log"
	"google.golang.org/grpc"
	"time"
)

func RpcLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx = GetReqidFromRpcContext(ctx) // 所有rpc请求进来，从client Context中获取的reqid, 设置到server context里,用于后续打印日志

	start := time.Now().UnixNano()
	resp, err := handler(ctx, req) // 真实的方法调用

	delay := (time.Now().UnixNano() - start) / int64(time.Millisecond)
	log.Infoc(ctx, fmt.Sprintf("Method:%v\nRequest:%v\nTime:%vms\nResponse:%v\n",
		info.FullMethod, req, delay, resp))
	return resp, err
}
