package main

import "fmt"

func main(){
	// 打印金字塔案例   
	// 1   2   3   4    5
	// var num int 
	// 
	// fmt.Printf("请输入金字塔层数:")
	// fmt.Scanln(&num)
	// for i :=1 ; i <= num ; i++ {
	// 	for j := 1 ; j <= i ; j++ {
	// 		fmt.Printf("*")
	// 	}
	// 	fmt.Printf("\n")
	// }

	// 1   3   5   7    9   居中
	// var num int 
	// var str01 string = "*"
	// fmt.Printf("请输入金字塔层数:")
	// fmt.Scanln(&num)
	// for i :=1 ; i <= num ; i++ {
	// 	// 在打印*前先打印空格
	// 	for k:= 1 ; k <= num - i ; k++ {
	// 		fmt.Printf(" ")
	// 	}
	// 	for j := 1 ; j <= 2*i - 1 ; j++ {
	// 		fmt.Printf(str01)
	// 	}
	// 	fmt.Printf("\n")
	// }

	// 第二种中间空心
	// 接收一个整数为层数
	var num int 
	var str01 string = "*"
	fmt.Printf("请输入金字塔层数:")
	fmt.Scanln(&num)
	for i :=1 ; i <= num ; i++ {
		// 在打印*前先打印空格
		for k:= 1 ; k <= num - i ; k++ {
			fmt.Printf(" ")
		}
		// j 表示每层打印多少
		for j := 1; j <= 2*i - 1; j++ {
			if j == 1 || j == 2 * i - 1 || i == num {
				fmt.Printf(str01)
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}

}