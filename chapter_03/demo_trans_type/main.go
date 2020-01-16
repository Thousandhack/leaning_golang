package main

import "fmt"
import "reflect"

func main(){
	// Goalng基本类型的数据转换
	var i int = 100
	// 将i转成float类型
	var n1 float32 = float32(i)
	fmt.Println("n1=",n1)
	fmt.Println(reflect.TypeOf(n1))
	
	var n2 int64 = int64(i)
	fmt.Println("n2=",n2)
	fmt.Println(reflect.TypeOf(n2))
}