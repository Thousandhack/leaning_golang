package main

import "fmt"

// 结构体嵌套

type address struct {
	province string
	city     string
}

type person struct {
	name string
	age int
	address  // 匿名嵌套结构体
}

type company struct{
	name string
	addr address
}

func main(){
	// 嵌套的结构体
	p1 := person{
		name: "zero",
		age:  18,
		address: address{
			province:"福建",
			city:"厦门",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.address.city)
	// 匿名嵌套可以直接查找出来，嵌套两个匿名结构体的话，就需要些全
	fmt.Println(p1.city)
}
