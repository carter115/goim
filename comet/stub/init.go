package stub

import (
	"context"
	"fmt"
	"gmimo/comet/config"
	"gmimo/common/constant"
	"gmimo/common/middleware"
	"google.golang.org/grpc"
)

var (
	cometAddr string
	logicAddr string
	pushAddr  string
)

func initAddr() {
	cometAddr = fmt.Sprintf("%s:%d", config.Rpc.CometSrv, config.Rpc.CometPort)
	logicAddr = fmt.Sprintf("%s:%d", config.Rpc.LogicSrv, config.Rpc.LogicPort)
	pushAddr = fmt.Sprintf("%s:%d", config.Rpc.PushSrv, config.Rpc.PushPort)
}

// 每一次rpc请求，都要加载请求上下文(设置包括超时时间，请求唯一reqid)
func rpcSetup(ctx context.Context) context.Context {
	rpcContext, _ := context.WithTimeout(context.Background(), config.Rpc.Timeout)

	// 设置超时，设置reqId
	reqid := ctx.Value(constant.REQID)
	if reqid != nil {
		rpcContext = middleware.SetReqidToRpcContext(reqid.(string), rpcContext)
	}

	// 程序启动，有可能配置文件还没加载成功
	if cometAddr == "" || logicAddr == "" || pushAddr == "" {
		initAddr()
	}
	return rpcContext
}

// 创建grpc client
func newLogicClient() (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	return grpc.Dial(logicAddr, opts...)
}
