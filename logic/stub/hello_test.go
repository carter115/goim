package stub

import (
	"context"
	helloProto "gmimo/common/proto/hello"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	addr := "logic.service:7001"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := helloProto.NewSayHelloClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rsp, err := client.Say(ctx, &helloProto.Request{Name: "Tom"})
	if err != nil {
		t.Fatal("client hello error: ", err)
	}
	t.Log("response:", rsp)

}
