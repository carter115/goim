package stub

import (
	"context"
	"gmimo/common/constant"
	messageProto "gmimo/common/proto/message"
	"google.golang.org/grpc"
	"testing"
)

func TestMessage(t *testing.T) {
	addr := "logic.service:7001"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := messageProto.NewSendMessageClient(conn)
	rsp, err := client.Send(context.Background(), &messageProto.Request{SrcId: "Tom", DstId: "ABCroom", MsgType: constant.P2Room})
	if err != nil {
		t.Fatal("client error: ", err)
	}
	t.Log("response:", rsp)
}
