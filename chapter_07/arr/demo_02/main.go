package main

import "fmt"

func main() {
	var intArr [4]int
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	intArr[3] = 40

	fmt.Println("第一个地址：", &intArr[0])
	fmt.Println("第二个地址：", &intArr[1])
	fmt.Println("第三个地址：", &intArr[2])
	fmt.Println("第四个地址：", &intArr[3])
	// 打印结果：
	// 第一个地址： 0xc000052120
	// 第二个地址： 0xc000052128
	// 第三个地址： 0xc000052130
	// 第四个地址： 0xc000052138

	// 遍历数组打印
	for i := 0; i < len(intArr); i++ {
		fmt.Println(intArr[i])
	}
}
