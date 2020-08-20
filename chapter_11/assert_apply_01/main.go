package main

import "fmt"


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
	b = a.(Point)
	fmt.Println(b)
}