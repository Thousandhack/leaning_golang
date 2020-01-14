package main

import "fmt"

func main(){
	// string 的基本使用
	var address string = "东北松花江"
	fmt.Println("address:",address)
	// 反引号
	var str1 string = `dqhodqi\n\n\n\n`
	fmt.Println("str1:",str1)
	// 字符串拼接方式
	var str2 string = "hello " + "world"
	fmt.Println("st2:",str2)
}