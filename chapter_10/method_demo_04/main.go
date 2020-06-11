package main

import "fmt"

type Circle struct {
	radius float64
}

// 为了提高效率，通常我们方法和结构体的指针类型绑定
func (c *Circle) area() float64 {
	// 这边的c 是指针，标准的访问方式为（*c）.radius
	fmt.Printf("c 是 *Circle只想的地址为： %p\n",c)
	(*c).radius = 5.0
	return 3.14 * (*c).radius * (*c).radius  // 这边也是使用指针  等价于：3.14 *c.radius * c.radius
}


func main(){
	//
	var c Circle
	fmt.Printf("main c 结构体变量的地址为：%p\n",&c)
	c.radius = 4.0

	res  := (&c).area()   //  返回值  地址* 值  也等价于 c.area()
	fmt.Println(res)
	fmt.Println((&c).radius)  // 值已经被改变，因为方法使用的是指针
}

/*
main c 结构体变量的地址为：0xc00000a0c0
c 是 *Circle只想的地址为： 0xc00000a0c0
78.5
5
*/