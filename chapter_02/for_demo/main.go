package main

import "fmt"

// for 循环

func main() {
	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 变种 1
	var j = 5
	for ; j < 10; j++ {
		fmt.Println(j)
	}

	// 变种 2
	var n = 5
	for n < 10 {
		fmt.Println("n=", n)
		n++
	}

	// 无限循环
	var m = 5
	for {
		if m > 10 {
			break
		}
		fmt.Println("m=", m)
		m++
	}

	// for range 循环
	s := "hello深圳"
	for k, v := range s {
		fmt.Printf("%d %c\n", k, v)
	}
}
