package main

import "fmt"

type Calcuator struct{

}


func (c Calcuator) NightNight(n int,m int) {
	for i:=1; i <= n; i++ {
		for j:=m; j >= i; j-- {
			fmt.Printf("%v * %v = %v ",i,j,i*j)
		}
		fmt.Println()
	}
	for q:=n; q >= 1; q-- {
		for a:=m; a >= q; a-- {
			fmt.Printf("%v * %v = %v ",q,a,q*a)
		}
		fmt.Println()
	}
}


func main(){
	// 打印出9*9乘法表
	var c Calcuator
	c.NightNight(9,9)
}