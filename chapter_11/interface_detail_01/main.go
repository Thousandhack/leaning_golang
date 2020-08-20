package main

import "fmt"

type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("stu say ")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i =", i)
}

type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Hello Monster")
}

func (m Monster) Say() {
	fmt.Println("Monster say...")
}

func main() {
	var stu Stu // 接口体变量实现了say方法，这个一定要
	var a AInterface = stu
	a.Say()

	var i integer = 10 // 不是结构体的变量也可以实现接口的应用
	var b AInterface = i
	b.Say()

	// 一个自定义类型可以实现多个接口
	// monster 实现AInterface 和 BInterface
	var monster Monster
	var a2 AInterface = monster
	var b2 BInterface = monster
	a2.Say()
	b2.Hello()


}
