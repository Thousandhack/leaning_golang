package main

import "fmt"

type MethodUtils struct {

}

func (mu MethodUtils) Print(){
	for i:=1; i <= 10; i++{
		for j:=1; j <= 8; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main(){
	// 编写结构体（MethodUtils）,编写一个方法，方法不需要参数
	// 在方法中打印10*8 的矩形，在main方法中调用
	var mu MethodUtils
	mu.Print()
}
