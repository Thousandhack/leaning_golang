package main

import "fmt"

// 构造函数
type person struct{
	name string
	age int
}

type dog struct {
	name string
}

// 返回结构体与结构体指针情况
// 当结构体比较大尽量使用结构体指针
// 约定俗成 new开头一般是构造函数
func newPerson(name string,age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}


func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func main(){
	p1 := newPerson("zero",18)
	p2 := newPerson("one",19)
	fmt.Println(p1,p2)
	dog1 := newDog("doudou")
	fmt.Println(dog1)
}
