package main

import "fmt"

type Monster struct{
	Name string
	Age int
}

type A struct{
	Monster
	int
}

func main(){
	// 使用匿名字段为基本数据类型的情况
	var a A
	a.Name = "test"
	a.Age = 10
	a.int = 80
	fmt.Println(a)
}