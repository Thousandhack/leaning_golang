package main

import "fmt"

func main(){
	// 用append 内置函数，可以对切片进行动态追加
	var slice []int = []int {100,200,300}
	fmt.Printf("slice=%v\n",slice)
	slice = append(slice,400)
	fmt.Println("new_slice=",slice)

	var demo_slice []int  = []int{1,2,3}
	// 切片追加切片
	slice = append(slice,demo_slice...)
	fmt.Println("new_new_slice=",slice)
	
}