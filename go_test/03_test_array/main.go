package main

import "fmt"

func arraySum() (sum int) {
	// 求数组[1 3 5 7 8] 所有元素的和
	array_demo := [...]int{1, 3, 5, 7, 8}
	//fmt.Println(sum)
	for _, value := range array_demo {
		sum += value
	}
	return sum
	//fmt.Println(sum)
}

func findSum() (a [2]int, b [2]int) {
	// 求一个数组和为8 的两个元素及所在位置
	array_demo := [...]int{1, 3, 5, 7, 8}
	//a := [...]int
	//b := [...]int
	for i := 0; i < len(array_demo); i++ {
		for j := 0; j < len(array_demo); j++ {
			if array_demo[i]+array_demo[j] == 8 {
				a[0] = i
				a[1] = array_demo[i]
				b[0] = j
				b[1] = array_demo[j]
				return a, b
			}

		}
	}
	return a,b
}

func main() {
	res := arraySum()
	fmt.Println("数组的和为:", res)
	res1, res2 := findSum()
	fmt.Println(res1)
	fmt.Println(res2)
}
