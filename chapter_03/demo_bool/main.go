package main

import "fmt"
import "unsafe"

func main(){
	// 注意事项
	// 1. bool类型占用存储空间是1个字节
	var b = false
	fmt.Println("b=",b)
	fmt.Println("b的占用空间=",unsafe.Sizeof(b))
	// 2.bool类型只能去true或者false
	
}