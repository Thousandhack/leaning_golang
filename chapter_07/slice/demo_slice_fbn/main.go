package main

import "fmt"

func fbn(n int) []uint64 {
	// 声明一个切片，切片大小n
	fbnSlice := make([]uint64, n)
	// 第一个数和第二个数的斐波那契为1
	if n == 1 {
		fbnSlice[0] = 1
	} else if n == 2 {
		fbnSlice[0] = 1
		fbnSlice[1] = 1
	} else {
		fbnSlice[0] = 1
		fbnSlice[1] = 1
		// 进行for循环来存放斐波那契的数列
		for i := 2; i < n; i++ {
			fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
		}
	}
	return fbnSlice
}

func main() {
	// 编写一个函数 fbn(n int )
	// 1. 可以接收一个 n int
	// 2.能够将斐波那契的数列放到切片中
	// 提示： slice[0] = 1  slice[1] = 1  slice[2] = 2   slice[3] = 3  slice[4] = 5
	fnbSlice := fbn(4)
	fmt.Println("fbnSlice=", fnbSlice)
}
