package main

import "fmt"

/*
一个景区根据年龄收取不同价格的门票，大于18收取  20    其他情况免费
编写Visitor结构体，根据年龄段决定能够购买门票的价格并输出

*/
type Visitor struct {
	name string
	age  int
}

func (v Visitor) Price() int {
	if v.age >= 18 {
		return 20
	} else {
		return 0
	}
}

func main() {
	var v Visitor
	//v.name = "zero"
	//v.age = 19
	for {
		fmt.Println("请输入你的名字：")
		fmt.Scanln(&v.name)
		if v.name == "n" {
			fmt.Println("退出程序...")
			break
		}
		fmt.Println("请输入你的年龄：")
		fmt.Scanln(&v.age)
		res := v.Price()
		fmt.Printf("%v的年龄为%v,门票价格为：%v元\n", v.name, v.age, res)
	}
	// 到 24分钟
}
