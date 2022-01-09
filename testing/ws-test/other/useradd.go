package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	wg     sync.WaitGroup
	prefix = "UserA"
	count  = 10000
	server = "http://mimo-logic01:10005"
)

func main() {
	wg.Add(count)
	for i := 0; i < count; i++ {
		uid := fmt.Sprintf("%s-%d", prefix, i)
		go createUser(&wg, uid)
	}
	wg.Wait()
}

// 用于批量创建用户
func createUser(wg *sync.WaitGroup, uid string) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	defer wg.Done()
	url := fmt.Sprintf("%s/user/generate?userId=%s", server, uid)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Response:", uid, string(result))
}
