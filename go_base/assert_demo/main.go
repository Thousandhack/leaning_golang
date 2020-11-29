package main

import "fmt"

// 类型断言


func main()  {
	var a interface{}  // 定义一个空接口
	a = 100
	// 如何判断a 保存的值得具体类型是什么？
	// 类型断言
	// x.(T)

	v1,ok := a.(int8)  // 这个地方不明白

	if !ok{
		fmt.Println("猜对了，a 是int",v1)
	} else {
		fmt.Println("猜错了")
	}
	// 2.switch
	fmt.Printf("%v\n",a)
	switch v2 := a.(type) {
	case int8:
		fmt.Println("猜对了，a 是int8",v2)
	case int16:
		fmt.Println("猜对了，a 是int16",v2)
	case int:
		fmt.Println("猜对了，a 是int",v2)
	default:
		fmt.Println("猜错了")
	}
}
