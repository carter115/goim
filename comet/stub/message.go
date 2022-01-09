package stub

import (
	"context"
	"gmimo/common/e"
	"gmimo/common/log"
	messageProto "gmimo/common/proto/message"
)

// rpc client

// 发送消息
func CallSendMessage(ctx context.Context, req *messageProto.Request) (rsp interface{}, errCode int) {
	rpcContext := rpcSetup(ctx) // 加载环境配置
	log.Debugcf(rpcContext, "[call Message] 消息ID:%s", req.Id)

	conn, err := newLogicClient() // 2. 根据配置，创建连接
	if err != nil {
		return e.RpcResponse(rpcContext, err)
	}

	client := messageProto.NewSendMessageClient(conn)
	_, err = client.Send(rpcContext, req)
	if err != nil {
		log.Errorcf(ctx, "[call Message error] %v", err)
		return e.RpcResponse(rpcContext, err)
	}

	return rsp, e.SUCCESS
}

// rpc service的方法
type MessageStub struct{}

// push服务推送消息到这里
func (msg *MessageStub) Send(ctx context.Context, req *messageProto.Request) (rsp *messageProto.Response, err error) {
	if ctx.Err() == context.Canceled {
		log.Warnc(ctx, e.GetMsg(e.ERR_RPC_TIMEOUT)+ctx.Err().Error())
		return nil, e.NewError(e.ERR_RPC_TIMEOUT)
	}

	log.Debugcf(ctx, "[recv Message] 消息ID:%v", req.Id)
	PushMessageChan <- req // 准备分发消息

	rsp = &messageProto.Response{}
	rsp.Status = e.GetMsg(e.SUCCESS)
	//fmt.Println("comet status:", rsp, req)
	return rsp, nil
}

// 消息分发通道
var PushMessageChan = make(chan *messageProto.Request, 1000)
