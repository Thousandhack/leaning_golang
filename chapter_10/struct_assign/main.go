package main

import "fmt"

type Stu struct {
	Name string
	Age  int
}

func main() {
	// 方式1：
	// 在创建结构体变量时，直接指定字段的值

	var stu1 = Stu{"hsz", 19}
	fmt.Println(stu1.Name, stu1.Age)
	stu2 := Stu{"zero", 20}
	fmt.Println(stu2.Name, stu2.Age)

	// 在结构体创建变量时，把字段和字段值写在一起
	var stu3 = Stu{
		Name: "jack",
		Age:  21,
	}
	fmt.Println(stu3.Name, stu3.Age)
	// Name 和 Age 可以更换位置
	stu4 := Stu{
		Name: "four",
		Age:  24,
	}
	fmt.Println(stu4.Name, stu4.Age)

	// 方式2 ：返回结构体的指针类型
	// var stu5 *Stu =
	var stu5 = &Stu{"five", 25}
	fmt.Println((*stu5).Name, (*stu5).Age)
	stu6 := &Stu{"six", 26}
	fmt.Println((*stu6).Name, (*stu6).Age)

	var stu7 = &Stu{
		Name: "seven",
		Age:  27,
	}
	fmt.Println((*stu7).Name, (*stu7).Age)

	stu8 := &Stu{
		Name: "eight",
		Age:  28,
	}
	fmt.Println(stu8.Name, stu8.Age) // 不加 * 号的方式也可以

}
