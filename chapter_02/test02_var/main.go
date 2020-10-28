package main

import "fmt"

// 声明变量
//var name string
//var age int
//var isOk bool

var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "hsz"
	age = 26
	isOk = true
	fmt.Println(name, age, isOk)
}
