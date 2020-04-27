package main

import "fmt"

func main(){
	// 二维数组例子
	/*
	0 0 0 0 0 0
	0 0 1 0 0 0
	0 2 0 3 0 0
	0 0 0 0 0 0
	*/
	// 声明二维数组
	var arr [4][6]int
	// 赋初值
	// fmt.Println(arr)

	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	// 遍历二维数组，按照要求输出图形
	for i :=0;i<4; i++ {
		for j :=0; j < 6; j++ {
			fmt.Print(arr[i][j]," ")
		}
		fmt.Println()
	}

}