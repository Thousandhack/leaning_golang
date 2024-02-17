package main

import "fmt"

func main() {
	// 指针
	// pointer
	// 1.& 取地址
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	// 2. * 根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	// 空指针不能赋值
	//var a *int // nil pointer
	var a = new(int) // new 申请一个内存赋值
	*a = 100
	fmt.Println(a) // 打印a的指针地址
}
