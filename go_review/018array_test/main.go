package main

import "fmt"

func main() {
	/*
		求数组[1, 3, 5, 7, 8]所有元素的和

	*/
	a1 := [...]int{1, 3, 5, 7, 8}
	total := 0
	for _, v := range a1 {
		total += v
	}
	fmt.Println(total)

	// 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	// 定义两个for循环，外层的从第一个开始遍历
	// 内存的for循环从外层后面的哪个开始找
	// 它们的两个数和为8
	for i := 0; i < len(a1); i++ {
		for j := 1; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Println(i, j)
			}

		}
	}

}
