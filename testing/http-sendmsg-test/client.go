package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

// gmimo提供发消息的http api，用于批量发送消息
func main() {
	var (
		url   = flag.String("url", "http://127.0.0.1:10000/message/send", "URL")
		count = flag.Int("count", 10, "请求数")
		wg    sync.WaitGroup
	)
	flag.Parse()
	wg.Add(*count)
	for i := 0; i < *count; i++ {
		msg := MessageBody{
			SrcId:   "su-test",
			DstId:   "abc",
			MsgType: "P2Room",
			Content: fmt.Sprintf("这是测试消息内容-%d", i),
		}
		bs, _ := json.Marshal(msg)
		go func(content []byte) {
			if _, err := http.Post(*url, "application/json", strings.NewReader(string(content))); err != nil {
				fmt.Println("error:", err)
			}
			wg.Done()
		}(bs)
	}
	wg.Wait()
	fmt.Println("执行完毕")
}

type MessageBody struct {
	SrcId   string `json:"srcId"`
	DstId   string `json:"dstId"`
	MsgType string `json:"msgType"`
	Content string `json:"content"`
}
