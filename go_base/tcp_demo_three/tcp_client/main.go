package main

import (
	"fmt"
	"gotest01/go_base/tcp_demo_three/proto"
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
	for i := 0; i < 20; i++ {
		msg := `Hello,How are you?`
		// 调用协议编码数据
		b,_ := proto.Encode(msg)
		conn.Write(b)

	}
}
