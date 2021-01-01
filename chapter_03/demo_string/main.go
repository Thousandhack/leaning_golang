package main

import (
	"fmt"
	"strings"
)

func main() {
	// string 的基本使用
	var address string = "东北松花江"
	fmt.Println("address:", address)
	// 反引号原样输出
	var str1 string = `dqhodqi\n\n\n\n`
	fmt.Println("str1:", str1)
	// 字符串拼接方式
	var str2 string = "hello " + "world"
	fmt.Println("st2:", str2)
	// 字符串相关操作
	fmt.Println("有多长:", len(str2))
	// 打印一个路径
	var path string = "D:\\goCode\\goproject\\src\\gotest01\\chapter_03\\demo_string"
	fmt.Println("path:", path)
	// 通过// 将字符串变量path进行分割形成一个数组
	res := strings.Split(path,"\\")
	fmt.Println(res)
	fmt.Printf("%T\n",res)
	// 多行的字符串
	s2 := `
	111111111111
	22222222222
	333333333
	`
	fmt.Println(s2)
}
