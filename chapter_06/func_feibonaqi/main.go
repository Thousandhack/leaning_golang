package main

import "fmt"

func fb(n int ) int {
	if (n == 2 || n == 1) {
		return 1
	} else {
		return fb(n-1) + fb(n-2)
	}
}


func main(){
	//
	num := fb(3)
	fmt.Println("num:",num)  // 输出第几个的斐波那契数
}
