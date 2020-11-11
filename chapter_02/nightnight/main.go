package main

import "fmt"
// 9*9 乘法表
func nightNight() {
	for i := 1; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Printf("%2v*%2v=%2v ", i, j, i*j)
		}
		fmt.Println("")
	}
}

func nightNightTwo(){
	fmt.Println("22222")
	for i := 1; i < 10; i++ {
		for j:=1; j < i+1; j++ {
			fmt.Printf("%2v*%2v=%2v ", i, j, i*j)
		}
		fmt.Println("")
	}
}

func main() {
	nightNight()
	nightNightTwo()
}
