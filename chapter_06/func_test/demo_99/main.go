package main

import "fmt"

func NightNight(){
	// 打印九九乘法表
	for i:=1;i<=9;i++{
		for j := 1; j <= i; j++{
			fmt.Printf("%v*%v=%v",i,j,i*j)
		}
		fmt.Println()
	}
}

func main(){
	NightNight()
}