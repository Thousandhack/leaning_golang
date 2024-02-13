package main

import "fmt"

func main() {
	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// 变种1
	var n = 5
	for ; n < 10; n++ {
		fmt.Println(n)
	}
	// 变种2
	var m = 5
	for m < 10 {
		fmt.Println(m)
		m++
	}

	// 无限循环
	//for {
	//	fmt.Println(11)
	//}

	s := "Hello, hsz"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
