package main

import "fmt"

// 单个声明常量
const pi = 3.1415926

// 批量声明常量
const (
	statusOK = 200
	notFund = 404
)

const (
	n1 = 100
	n2
	n3
)


func main(){
	fmt.Println(pi)
	fmt.Println(statusOK)
	fmt.Println(notFund)
	fmt.Println(n1,n2,n3)
}