package main


import "fmt"

type Person struct {
	Name string
}

// 将相当于给Person 绑定了一个方法
func (p Person) test() { 
	fmt.Println("Name=", p.Name)
}

func main() {
	var p Person
	p.Name = "hsz"
	p.test()  // 调用方法
}


/*
说明：
test方法和Person类型绑定
test方法只能通过Person类型的变量来调用，不能直接调用,也不能使用其他类型的变量来调用
func (p Person)  表示哪个Person 变量调用，这个p就是哪个副本
*/