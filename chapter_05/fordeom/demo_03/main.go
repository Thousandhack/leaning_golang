package main

import "fmt"

func main(){
	// 打印1~100之间倍数是9的整数的个数及总和
	var count int = 0
	var sum int = 0
	for i:=1 ; i <= 100 ;i ++ {
		if i % 9 == 0 {
			sum += i
			count ++
		}
	}
	fmt.Println("个数为：",num)
	fmt.Println("总和为：",sum)
}