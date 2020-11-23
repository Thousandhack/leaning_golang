package main

import "fmt"

// 接口的实现

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("爬房梁")
}

func (c cat) eat(food string) {
	fmt.Printf("%s吃%s~\n",c.name,food)
}

func (c chicken) move() {
	fmt.Println("挥动翅膀")
}

func (c chicken) eat(food string) {
	fmt.Printf("%s吃%s~\n",c.name,food)
}

func main() {
	var a1 animal  // 最开始的动态类型为nil 动态值为nil

	cat1 := cat{ // 定义一个cat类型的变量
		name: "小黑猫",
		feet: 4,
	}
	a1 = cat1  // 接口：动态类型和动态值  接口保存的是：一部分存赋值的类型，一部分存值
	// 实现接口变量能够实现不同的值
	a1.move()
	a1.eat("鱼")

	c1 := chicken{ // 定义一个cat类型的变量
		name: "小公鸡",
		feet: 4,
	}
	a1 = c1
	c1.move()
	c1.eat("青菜虫")
}
