package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// UDP client

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 40000})
	if err != nil {
		fmt.Println("connect failed,err:", err)
		return
	}
	defer socket.Close()
	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入内容：")
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg))
		// 收到回复的数据
		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("redv reply msg failed,err:", err)
		}
		fmt.Println("收到回复：", string(reply[:n]))
	}

}
