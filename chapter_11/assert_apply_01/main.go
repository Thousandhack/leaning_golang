package main

import (
	"fmt"
)


type Point struct {
	x int
	y int
}

func main(){
	var a interface{}
	var point Point = Point{1,2}
	a = point
	var b Point
	// 将一个空接口类型转换成对应得类型，使用类型断言
	// a.(Point) 表示就是类型断言，表示判断a是否指向Point类型的变量，如果是就转成Point类型并赋给b变量，否则保错
	b = a.(Point)
	fmt.Println(b)

	// 类型断言（带检测的）
	var x interface{}
	var b2 float32 = 1.1
	x = b2 // 空接口，可以接收任意类型
	// x = > float32 [使用类型断言]

	y,ok := x.(float32)   // 这样的情况下能检查断言成功还是失败
	if ok == true{
		fmt.Println("convert success")
		fmt.Printf("y的类型是 %T 值是=%v",y,y)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行...")
}