package main

import "fmt"

type Circle struct {
	radius float64
}

// 声明一个圆面积与半径的关系
func (c Circle) area() float64 { // 可以改为指针的方式： Circle  -- 》 *Circle
	return 3.14 * c.radius * c.radius
}

func main() {
	// 声明一个circle，字段有radiusd的结构体
	// 声明一个方法area和Circle绑定，返回面积
	//
	var c Circle
	c.radius = 4.0 // 指定半径
	theArea := c.area()
	fmt.Println("the area is :", theArea)

}
