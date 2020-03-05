package main

import (
	"fmt"
)

func test(n int) {
	if n > 2 {
		n--
		test(n) // 把test()函数全部带入
	}
	fmt.Println("n=", n)  // n = 2 n =2  n = 3
}

func test2(n int) {
	if n > 2 { // 注意执行过if 不在执行else
		n--
		test(n) // 把test()函数全部带入
	} else {
		fmt.Println("n=", n)  //  n = 2
	}	
}

func main() {
	test(4)

	test2(4)   
}
