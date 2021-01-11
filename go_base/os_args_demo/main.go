package main

import (
	"fmt"
	"os"
)

// 运行命令： go run main.go 1 2 3
func main() {
	fmt.Println(os.Args)
	fmt.Printf("%T\n", os.Args) // 打印数据类型
	fmt.Println("123456789")
}
