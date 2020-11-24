package main

import "fmt"

// 接口与类型的关系
// 多个类型可以实现同一个接口
// 一个类型（同一个结构体）可以实现多个接口

// 接口嵌套
type animal interface {
	move()  // mover
	eat(string)   // eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

// cat结构体实现mover和eater接口
type cat struct {
	name string
	feet int8
}


func (c cat)eat(food string){
	fmt.Printf("猫吃..%s\n",food)
}


func (c cat)move(){
	fmt.Println("在跑...")
}


func main()  {
	var m mover
	c1 := cat{
		name: "小白猫",
		feet: 4,
	}
	m = c1
	m.move()

	var e eater
	e = c1
	e.eat("小鱼")

}
