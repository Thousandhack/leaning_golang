package main

import "fmt"

// 一个test函数
func test(n1 int){
	n1 += 1
	fmt.Println("n1=",n1)
}

func main(){
	// 调用 函数
	n1 := 10
	test(n1)
	fmt.Println("n1=",n1)

}