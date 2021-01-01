package main

import "fmt"

// 给自定义类型加方法

type  myInt int

func (m myInt)hello()  {
	fmt.Println("我是一个自己的int类型")
	fmt.Printf("%T\n",m)
}

func main()  {
	// myInt 类型的值为100
	m := myInt(100)
	m.hello()

	//
	var i int32 = 10
	fmt.Println(i)
	var o = int32(10)
	fmt.Println(o)
}
