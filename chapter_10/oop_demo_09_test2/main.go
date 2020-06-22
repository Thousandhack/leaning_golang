package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOK", a.Name)
}

func (a *A) Hello() {
	fmt.Println("A Hello", a.Name)
}

type B struct {
	A
	Name string
}

// 建立B的两个方法
func (b *B) SayOk() {
	fmt.Println("B SayOk", b.Name)
}


// func (b *B) Hello() {
// 	fmt.Println("B Hello", b.Name)
// }

func main() {

	var b B
	b.A.Name = "zero"
	b.A.age = 10
	b.A.SayOk()
	b.A.Hello()

	// 上面的写法可以简化
	b.Name = "one"
	b.SayOk()
	b.Hello()   // 因为b 中没有Hello方法，所以自动继承a中的Hello方法
	// 上面会先到b 去寻找是否有这个字段，
	// 再到它继承的A中查询是否有这个字段
}
