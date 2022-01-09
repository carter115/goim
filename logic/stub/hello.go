package stub

import (
	"context"
	"gmimo/common/e"
	"gmimo/common/log"
	helloProto "gmimo/common/proto/hello"
)

type HelloStub struct{}

func (s *HelloStub) Say(ctx context.Context, req *helloProto.Request) (rsp *helloProto.Response, err error) {
	if ctx.Err() == context.Canceled {
		log.Warnc(ctx, e.GetMsg(e.ERR_RPC_TIMEOUT)+ctx.Err().Error())
		return nil, e.NewError(e.ERR_RPC_TIMEOUT)
	}

	log.Debugcf(ctx, "[recv Hello] %v", req)
	rsp = &helloProto.Response{}
	rsp.Msg = "Hi " + req.Name
	return rsp, nil
}
