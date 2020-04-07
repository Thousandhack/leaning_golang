package main

import "fmt"

func main() {
	//string底层是一个byte数组 ，因此string也可以进行切片
	str := "hello_hello_hsz"
	// 使用切片获取hsz
	slice := str[12:]
	fmt.Println("slice=", slice)

	// string --> byte  转成byte数组
	arr := []byte(str)
	arr[0] = 'z'
	str = string(arr)
	fmt.Println("str=", str)

	// 可以处理英文和数字，但是不能处理中文
	// []byte 字节来处理，而一个汉字，是3个字节，因此就会出现乱码
	// 解决方法： string 转成 []rune 因此rune是按字符来计算处理的，兼容汉字
	arr1 := []rune(str)
	arr1[0] = '好'
	str = string(arr1)
	fmt.Println("str=", str)
}
