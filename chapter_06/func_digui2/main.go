package main

import "fmt"

// func test(n int) int {
// 	if n == 1 {
// 		return 3
// 	} else {
// 		return 2 * test(n-1) + 1
// 	}
// }


func test(n int) int {
	if n == 10 {
		return 1
	} else {
		return 2 * (test(n+1) + 1)
	}
}



func main(){
	// fmt.Println("f(1)=",test(1))  // 3
	// fmt.Println("f(5)=",test(5))  // 63
	fmt.Println(test(8))
}

