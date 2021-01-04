package main

import (
	"fmt"
	"net"
)
// 客户端
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	// 这边的一条信息，发了20次，太快了造成黏包
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		//time.Sleep(time.Second)  // 这样加时间就不会
		conn.Write([]byte(msg))


	}
}
