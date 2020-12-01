package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取用户输入时有空格

func useScan(){
	var s string
	fmt.Print("请输入内容:")
	fmt.Scanln(&s)  // 这种输入只能匹配连续的字符中间不能出现空格
	fmt.Printf("你输入的内容是:%s\n",s)
}

func useBufio(){
	var s string
	fmt.Print("请输入内容:")
	reader := bufio.NewReader(os.Stdin)
	s,_ = reader.ReadString('\n')  // 读取换行前的数据
	fmt.Printf("你输入的内容是:%s\n",s)

}

func main()  {
	// useScan()
	useBufio()
}
