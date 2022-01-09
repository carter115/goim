package stub

import (
	"context"
	"gmimo/common/e"
	"gmimo/common/log"
	roomProto "gmimo/common/proto/room"
)

func CallRoomJoin(ctx context.Context, req *roomProto.Request) (rsp interface{}, errCode int) {
	rpcContext := rpcSetup(ctx) // 1. 装载环境配置
	log.Debugcf(rpcContext, "[call RoomJoin] %+v", req)

	conn, err := newLogicClient() // 2. 根据配置，创建连接
	if err != nil {
		return e.RpcResponse(rpcContext, err)
	}

	// rpc调用logic
	cli := roomProto.NewRoomClient(conn)
	if _, err = cli.Join(rpcContext, req); err != nil {
		return e.RpcResponse(rpcContext, err)
	}
	return
}

func CallRoomLeave(ctx context.Context, req *roomProto.Request) (rsp interface{}, errCode int) {
	rpcContext := rpcSetup(ctx)
	log.Infocf(rpcContext, "[call RoomLeave] %+v", req)

	conn, err := newLogicClient()
	if err != nil {
		return e.RpcResponse(rpcContext, err)
	}

	// rpc调用logic
	cli := roomProto.NewRoomClient(conn)
	if _, err = cli.Leave(rpcContext, req); err != nil {
		return e.RpcResponse(rpcContext, err)
	}

	return
}

func CallRoomMember(ctx context.Context, req *roomProto.Request) (rsp interface{}, errCode int) {
	rpcContext := rpcSetup(ctx)
	log.Infocf(rpcContext, "[call RoomMember] %+v", req)
	conn, err := newLogicClient()
	if err != nil {
		return e.RpcResponse(rpcContext, err)
	}
	// rpc调用logic
	cli := roomProto.NewRoomClient(conn)
	if rsp, err = cli.Member(rpcContext, req); err != nil {
		return e.RpcResponse(rpcContext, err)
	}
	return
}
