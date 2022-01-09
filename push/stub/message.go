package stub

import (
	"context"
	"gmimo/common/e"
	"gmimo/common/log"
	messageProto "gmimo/common/proto/message"
)

// TODO RPC连接所有comet，把一条消息推送到所有comet

// rpc client

// 发送消息
func CallSendMessage(ctx context.Context, req *messageProto.Request) (rsp interface{}, errCode int) {
	ctx = rpcSetup(ctx) // 加载环境配置
	log.Debugcf(ctx, "[call Message] 消息ID:%s", req.Id)

	conn, err := newCometClient() // 2. 根据配置，创建连接
	if err != nil {
		log.Errorcf(ctx, "[new comet rpc client error] %v", err)
		return e.RpcResponse(ctx, err)
	}

	client := messageProto.NewSendMessageClient(conn)
	rsp, err = client.Send(ctx, req)
	if err != nil {
		log.Errorcf(ctx, "[call Message error] %v", err)
		return e.RpcResponse(ctx, err)
	}

	return rsp, e.SUCCESS
}
