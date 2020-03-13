package main

import "fmt"
import "strings"

func makeSuffix(suffix string) func(string) string {
	// （1）编写一个函数 makeSuffix (suffix string) 可以接收一个文件后缀名比如后缀名(.jpg)
	// 并返回一个闭包
	// （2）调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀比如后缀名(.jpg)，则返回
	// 文件名.jpg,如果已经有.jpg后缀，则返回原文件名

	// （3） 要求使用闭包的方式完成
	// （4）strings.HasSuffix  这个函数可以判断是否存有指定的后缀
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	// .jpg的后缀
	f := makeSuffix(".jpg")
	// 实例引用
	fmt.Println("new_name:", f("today"))

	fmt.Println("new_name:", f("today.jpg"))
	

}

// 以上代码说明：返回的匿名函数和maksSuffix(suffix string)的suffix 变量组成一个闭包
// 返回的函数引用到suffix变量
