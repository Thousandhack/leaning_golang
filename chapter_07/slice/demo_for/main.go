package main

import "fmt"

func main() {
	// 使用常规的for循环遍历切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	myslice := arr[1:4] // 第二个元素到第4个元素
	for i := 0; i < len(myslice); i++ {
		fmt.Printf("i=%v v=%v\n",i,myslice[i])
	}

	// 使用for--range方式遍历切片
	for i, v := range myslice {
		fmt.Printf("i=%v v=%v\n", i, v)
	}

	// 切片还可以切片
	slice2 := myslice[1:]
	fmt.Println(slice2)
}
