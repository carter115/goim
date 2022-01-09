package e

import (
	"context"
	"gmimo/common/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 处理client调用的出错
func RpcResponse(ctx context.Context, err error) (rsp interface{}, errCode int) {
	statusErr, ok := status.FromError(err)
	if ok {
		switch statusErr.Code() {
		case codes.DeadlineExceeded:
			errCode = ERR_RPC_TIMEOUT // 请求超时
		default:
			errCode = ERR_RPC_ERROR // 网络不可达或者其它未定义错误
		}
		rsp = err.Error()
		log.Warncf(ctx, "%s:%v", GetMsg(errCode), rsp)
	}
	return
}
