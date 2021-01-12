package main

import (
	"flag"
	"fmt"
)

// flag 获取命令参数
func main(){
	// 创建一个标志位参数
	name := flag.String("name","hsz","请输入名字")
	// 使用flag
	flag.Parse()
	fmt.Println(*name)
}
