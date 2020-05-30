package main

import "fmt"

type A struct{
	Num int
}

type B struct{
	Num int
}

func main()  {
	// 如果要在两个结构体间转换必须两个结构体字段相同
	// 字段个数，字段名称，字段类型相同
	var a A
	var b B
	a = A(b)
	fmt.Println(a,b)

}