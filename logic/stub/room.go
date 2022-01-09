package stub

import (
	"context"
	"gmimo/common/e"
	"gmimo/common/log"
	roomProto "gmimo/common/proto/room"
	"gmimo/logic/controller"
)

// rpc server的方法
type RoomStub struct{}

func (*RoomStub) Join(ctx context.Context, req *roomProto.Request) (rsp *roomProto.Response, err error) {
	if ctx.Err() == context.Canceled {
		log.Warnc(ctx, e.GetMsg(e.ERR_RPC_TIMEOUT)+ctx.Err().Error())
		return nil, e.NewError(e.ERR_RPC_TIMEOUT)
	}
	log.Debugcf(ctx, "[recv RoomJoin] %v", req)
	rsp = &roomProto.Response{Status: e.GetMsg(e.SUCCESS)}

	m := &controller.Room{req.Mid}
	if err := m.Join(req.Uid); err != nil {
		rsp.Status = e.GetMsg(e.ERR_ROOM_JOIN)
	}
	return rsp, nil
}

func (*RoomStub) Leave(ctx context.Context, req *roomProto.Request) (rsp *roomProto.Response, err error) {
	if ctx.Err() == context.Canceled {
		log.Warnc(ctx, e.GetMsg(e.ERR_RPC_TIMEOUT)+ctx.Err().Error())
		return nil, e.NewError(e.ERR_RPC_TIMEOUT)
	}

	log.Debugcf(ctx, "[recv RoomLeave] %v", req)
	rsp = &roomProto.Response{Status: e.GetMsg(e.SUCCESS)}

	m := &controller.Room{Id: req.Mid}
	m.Leave(req.Uid) // 离开房间功能，忽略离开失败
	return rsp, nil
}

func (*RoomStub) Member(ctx context.Context, req *roomProto.Request) (rsp *roomProto.RespMember, err error) {
	if ctx.Err() == context.Canceled {
		log.Warnc(ctx, e.GetMsg(e.ERR_RPC_TIMEOUT)+ctx.Err().Error())
		return nil, e.NewError(e.ERR_RPC_TIMEOUT)
	}
	log.Debugcf(ctx, "[recv RoomMember] %v", req)
	rsp = &roomProto.RespMember{Status: e.GetMsg(e.SUCCESS)}

	m := &controller.Room{Id: req.Mid}
	rsp.Member, _ = m.Member()
	return rsp, nil
}
