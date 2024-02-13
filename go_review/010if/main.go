package main

import "fmt"

func main() {
	// if 条件判断
	age := 14
	if age > 18 {
		fmt.Println("可以进澳门赌场")
	} else {
		fmt.Println("不能去，需要写作业，好好读书")
	}

	if age > 35 && age < 60 {
		fmt.Println("中年")
	} else if age > 60 {
		fmt.Println("老年")
	} else {
		fmt.Println("青少年")
	}
}
