package main

import "fmt"

// 结构体嵌套

type address struct {
	province string
	city     string
}

type person struct {
	name    string
	age     int
	address // 匿名嵌套结构体
}

type company struct {
	name string
	addr address
}

// 结构体模拟实现其他语言中的"继承"
type animal struct {
	name string
}

// 给 animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

// 狗类
type dog struct {
	feet uint8
	animal
}

// 给dog实现一个汪汪汪的方法
func (d dog) wang(){
	fmt.Printf("%s汪汪汪\n",d.animal.name)
}

func main() {
	// 嵌套的结构体
	p1 := person{
		name: "zero",
		age:  18,
		address: address{
			province: "福建",
			city:     "厦门",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.address.city)
	// 匿名嵌套可以直接查找出来，嵌套两个匿名结构体的话，就需要些全
	fmt.Println(p1.city)

	// 实现结构体继承结构体
	d1 := dog{
		feet:   4,
		animal: animal{
			name:"doudou",
		},
	}
	d1.move()
	d1.wang()
}
