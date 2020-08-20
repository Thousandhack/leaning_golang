package main

import "fmt"

type CInterface interface{
	test01()
}

type BInterface interface{
	test02()
}

type AInterface interface{
	CInterface
	BInterface
	test03()
}

// 如果需要实现A接口，就需要是西安B，C接口

type Stu struct {

}

func (stu Stu) test01(){
	fmt.Println("test01")
}

func (stu Stu) test02(){
	fmt.Println("test02")
}

func (stu Stu) test03(){
	fmt.Println("test03")
}


// 定义空接口的第一种方式
type T interface {

}

func main(){
	var stu Stu
	var a AInterface = stu
	a.test01()
	a.test02()
	a.test03()

	// 空接口
	var t T = stu
	fmt.Println(t)  // 就是打印空接口

	// 定义空接口的第二后种方式
	var t2 interface{} = stu
	var num1 float64 = 8.8
	// 把float64的值赋给空接口
	t2 = num1
	fmt.Println(t2)
}
