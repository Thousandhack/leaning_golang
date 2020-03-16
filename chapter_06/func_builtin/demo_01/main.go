package main

import "fmt"

func main() {
	//
	num1 := 100
	fmt.Printf("num1的类型%T,num1的值=%v,num1的地址%v\n",num1,num1,&num1)

	num2 := new(int)  // num2是指针
	// 指针的值是一个地址
	// num2 的类型%T => *int
	// num2 的值 = 地址 0xc0000540a0
	// num2的地址 = 0xc000080020
	
	// 修改num2 指向的的值
	*num2 = 100
	fmt.Printf("num2的类型%T,num2的值=%v,num2的地址%v 指向的值=%v\n",num2,num2,&num2,*num2)

}
