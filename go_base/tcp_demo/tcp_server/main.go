package main

import (
	"fmt"
	"net"
)

// tcp server 端

func processConn(conn net.Conn){
	// 3. 与客户通信
	var tmp [128]byte
	n, err := conn.Read(tmp[:])
	if err != nil {
		fmt.Println("read err")
		return
	}
	fmt.Println(string(tmp[:n]))
}

func main() {
	// 1.本地端口启动服务
	listen, err := net.Listen("tcp", "127.0.0.1:19999")
	if err != nil {
		fmt.Println("start server failed")
		return
	}
	for {
		// 2.等待连接建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed")
			return
		}
		go processConn(conn)
	}
}
