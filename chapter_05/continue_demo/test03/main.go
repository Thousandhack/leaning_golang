package main

import "fmt"

func main() {
	// 输入个数不确定的整数，并判断读入的整数和负数的个数
	// 输入为0时结束程序
	// positive 正数  negitive 负数
	// var positive int 
	var p_count int = 0
	// var negitive int
	var n_count int = 0
	var num int
	for {
		fmt.Println("请输入一个整数:")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}

		if num > 0 {
			p_count ++
			continue
		} 
		n_count ++

	}
	fmt.Printf("正数个数是%v",p_count)
	fmt.Printf("负数个数是%v",n_count)
}