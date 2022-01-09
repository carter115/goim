package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"gmimo/common/constant"
	"gmimo/common/util"
	"google.golang.org/grpc/metadata"
)

var reqKey = constant.REQID

// 设定请求ID
func SetRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(reqKey, util.GetUUID())
		c.Next()
	}
}

// rpc client: 设置reqid的Context，发送给服务端
func SetReqidToRpcContext(reqid string, src context.Context) context.Context {
	src = context.WithValue(src, reqKey, reqid) // reqid用于后续打印日志
	md := metadata.Pairs(reqKey, reqid)         // reqid用于后续发送给服务端
	return metadata.NewOutgoingContext(src, md)
}

// rpc server: 从客户端的Context提取Context中的reqid
func GetReqidFromRpcContext(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx) // 从client传过来的rpc Context获取数据
	if ok && len(md[reqKey]) > 0 {
		ctx = context.WithValue(ctx, reqKey, md[reqKey][0])
	}
	return ctx
}
