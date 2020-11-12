package main

import (
	"fmt"
	"time"
)

// 9*9 乘法表
func nightNightOne() {
	for i := 1; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Printf("%2v*%2v=%2v ", i, j, i*j)
		}
		fmt.Println("")
	}
}

func nightNightTwo() {
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%2v*%2v=%2v ", i, j, i*j)
		}
		fmt.Println("")
	}
}

func nightNightThree() {
	for i := 9; i >= 1; i-- {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%2v*%2v=%2v ", i, j, i*j)
		}
		fmt.Println("")
	}
}

func nightNightFour() {
	for i := 9; i >= 1; i-- {
		for j := 9; j >= i; j-- {
			fmt.Printf("%2v*%2v=%2v ", i, j, i*j)
		}
		fmt.Println("")
	}
}

func main() {
	nightNightOne()
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("    ================================      ")
	fmt.Println("===========================================")
	nightNightTwo()
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("===========================================")
	nightNightThree()
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("============================================")
	nightNightFour()
}
