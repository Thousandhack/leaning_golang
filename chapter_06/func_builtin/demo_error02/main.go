package main

import (
	"errors"
	"fmt"
)

// 函数去读取以配置文件init.conf的信息
// 如果文件名传入不正确，我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		// 读取
		return nil
	} else {
		// 返回一个自定义错误
		return errors.New("读取文件错误...")
	}
}

func test() {
	err := readConf("config.ini2") // 看传的文件名是否正确，不正确就报错
	if err != nil {
		// 如果读取文件发送错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test继续执行")
}

func main() {
	// 自定义错误的使用
	test()
	fmt.Println("运行main()后面的代码")
}
