package connection

import (
	"fmt"
	"log"
	"testing"
)

func TestPubsub(t *testing.T) {
	if err := Publish(UserConnEvent{"jack", "0xjabclk", LeaveRoom}); err != nil {
		log.Fatalln("publish error:", err)
	}
	fmt.Println("send message successful.")

}
