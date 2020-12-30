package main

import (
	"fmt"
	"net"
)

// tcp client
func main() {
	// 1.与server 端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:19999")
	if err != nil {
		fmt.Println("dial failed")
		return
	}
	// 2.发送数据
	conn.Write([]byte("hello hsz!"))
	conn.Close()
}
