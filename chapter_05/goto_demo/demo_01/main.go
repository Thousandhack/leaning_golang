package main

import "fmt"

func main() {
	// 演示goto的使用
	var n int = 25
	fmt.Println("ok1")
	fmt.Println("ok2")
	if n > 20 {
		goto label1
	}
	fmt.Println("ok3")
	fmt.Println("ok4")
	fmt.Println("ok5")
	label1:
	fmt.Println("ok6")
	fmt.Println("ok7")
}