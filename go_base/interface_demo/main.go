package main

import "fmt"

// 引出接口的实例

// 定义一个能叫的类型
type speaker interface {
	speak()
}
type cat struct {
	name string
}

type dog struct {
	name string
}

type person struct {
	name string
}

func da(x speaker) {
	// 接收一个参数(不管什么类型)，传进来什么，我就打什么
	x.speak()
}

func (c cat) speak() {
	fmt.Printf("%s 会叫：喵喵喵~\n", c.name)
}

func (d dog) speak() {
	fmt.Printf("%s 会叫：汪汪汪~\n", d.name)
}

func (p person) speak() {
	fmt.Printf("%s 会叫：啊啊啊~\n", p.name)
}

func main() {
	var c1 = cat{
		name: "mimi",
	}
	var d1 = dog{
		name: "doudou",
	}
	var p1 = person{
		name: "mingming",
	}
	da(c1)
	da(d1)
	da(p1)
}
