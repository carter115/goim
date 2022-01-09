package stub

import (
	"context"
	"gmimo/common/e"
	"gmimo/common/log"
	helloProto "gmimo/common/proto/hello"
)

// rpc client
func CallHello(ctx context.Context, name string) (rsp interface{}, errCode int) {
	rpcContext := rpcSetup(ctx) // 1. 加载环境配置
	log.Infocf(ctx, "[call Hello] %s", name)

	conn, err := newLogicClient() // 2. 根据配置，创建连接
	if err != nil {
		return e.RpcResponse(rpcContext, err)
	}

	client := helloProto.NewSayHelloClient(conn)
	req := &helloProto.Request{Name: name}
	_, err = client.Say(rpcContext, req)
	if err != nil {
		return e.RpcResponse(rpcContext, err)
	}
	return rsp, e.SUCCESS
}
