package main

import "fmt"

func main(){

	// 创建一个byte类型的26个元素的数组，分别放置A-Z。使用for循环访问所有元素并打印
	// 出来。提示：字符数据运算 A + 1 > B

	// 思路：
	//1.声明一个数组 var myChars [26]bytes
	// 2.使用for循环，利用 字符可以进行运算的特点来赋值  'A' + 1 -> 'B'

	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i)   // 注意需要将 i = > byte
	}

	for i:=0; i < 26; i++ {
		fmt.Printf("%c ",myChars[i])
	}
}