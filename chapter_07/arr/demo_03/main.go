package main

import "fmt"

func main() {
	// 四种数组的初始化
	// 法一
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01", numArr01)

	// 法二
	var numArr02 = [3]int{4, 5, 6}
	fmt.Println("numArr02", numArr02)

	// 法三
	var numArr03 = [...]int{7, 8, 9}
	fmt.Println("numArr03", numArr03)

	// 法四
	var numArr04 = [...]int{1:11, 0:10, 2:12}  // 按下表的位置顺序输出，有序
	fmt.Println("numArr04", numArr04)


	var arrDemo = [...]int{1,2,3,4,5}
	// for-range遍历方式
	for index, value := range arrDemo{
		fmt.Println(index,value)
	}
}
