package main

import "fmt"

type Person struct{
	Name string
	Age int
}

func main(){
	// 以下p1 和 p2 共享一个内存地址
	var p1 Person
	p1.Age = 10
	p1.Name = "小明"
	var p2 *Person = &p1   // p2是一个指针

	fmt.Println((*p2).Age)  //golang中 (*p2) 与 p2相同
	fmt.Println(p2.Age)

	p2.Name = "SIX"
	fmt.Println(p2.Name)
	fmt.Println(p1.Name)
	fmt.Println((*p2).Name)
}