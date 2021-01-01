package main

import "fmt"

func main() {
	// 浮点类型的使用
	var price float32 = 19.23
	fmt.Println("price:", price)

	// 尾数部分可能丢失，造成精度损失
	var num3 float32 = -123.2345623456
	var num4 float64 = -123.2345623456
	fmt.Println("num3:", num3)
	fmt.Println("num4:", num4)
	// 打印结果如下：
	// num3: -123.234566
	// num4: -123.2345623456

	// 测试浮点型的默认精度
	var b = 1.11
	fmt.Printf("b的类型为%T\n", b)

	// 科学计算法
	num5 := 5.12e2
	fmt.Println("num5:", num5)

	// 基础数据的数据类型的默认值
	var q int 
	var w float32
	var e float64
	var isMarryied bool
	var name string
	fmt.Println("q=",q)
	fmt.Println("w=",w)
	fmt.Println("e=",e)
	fmt.Println("isMarryied=",isMarryied)
	fmt.Println("name=",name)

	// 小数默认float64
	// 不同类型不能进行赋值
	f1 := 1.2345678
	fmt.Printf("%T\n",f1)
}
