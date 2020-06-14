package main

import "fmt"


type Person struct{
	Name string
}

func test1(p Person){
	fmt.Println(p.Name)
}

func test2(p *Person){
	fmt.Println(p.Name)
}

func main() {
	// 函数
	// 对于普通函数，接受者为值类型时，不能将指针类型的数据直接传递，反之也不行
	p := Person{"zero"}
	test1(p)  // 值类型传递值
	test2(&p) // 指针类型传递地址
}