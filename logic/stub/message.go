package stub

import (
	"context"
	"encoding/json"
	"gmimo/common/e"
	"gmimo/common/log"
	messageProto "gmimo/common/proto/message"
	"gmimo/logic/controller"
)

// rpc service的方法
type MessageStub struct{}

func (msg *MessageStub) Send(ctx context.Context, req *messageProto.Request) (rsp *messageProto.Response, err error) {
	if ctx.Err() == context.Canceled {
		log.Warnc(ctx, e.GetMsg(e.ERR_RPC_TIMEOUT)+ctx.Err().Error())
		return nil, e.NewError(e.ERR_RPC_TIMEOUT)
	}

	log.Debugcf(ctx, "[recv Message] 消息ID:%v", req.Id)
	rsp = &messageProto.Response{}
	data, _ := json.Marshal(req)
	if err := controller.KafkaPush(ctx, data); err != nil { // 推送到kafka
		rsp.Status = e.GetMsg(e.ERR_KAFKA_PUSH_MSG)
		return rsp, err
	}

	rsp.Status = e.GetMsg(e.SUCCESS)
	return rsp, nil
}
