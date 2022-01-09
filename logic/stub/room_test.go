package stub

import (
	"context"
	roomProto "gmimo/common/proto/room"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestRoom(t *testing.T) {
	addr := "logic.service:7001"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := roomProto.NewRoomClient(conn)
	req := &roomProto.Request{Mid: "XYZroom", Uid: "Jack"}

	// 加入房间
	rsp, err := client.Join(context.Background(), req)
	if err != nil {
		t.Fatal("client error: ", err)
	}
	t.Log("join response:", rsp)

	// 10秒后离开房间
	time.Sleep(10 * time.Second)
	rsp, err = client.Leave(context.Background(), req)
	if err != nil {
		t.Fatal("client error: ", err)
	}
	t.Log("leave response:", rsp)
}
