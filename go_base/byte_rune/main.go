package main

import "fmt"

func main() {
	s := "Hello_hsz"
	n := len(s)
	fmt.Println(n)
	// byte类型和rune 类型
	// go语言中为了处理非ASCII码类型的字符 定义了新的rune类型
	// rune 类型代表utf-8字符
	for _, c := range s {
		//fmt.Println(c)
		// 需要进行格式化输出，否则会打印字符byte
		fmt.Printf("%c\n", c)
	}


	// 字符串修改（本身不能修改）
	// 转换成另外一种变量
	s1 := "白萝卜"    // 双引号表示字符
	s2 := []rune(s1)  // 把字符串强制转成一个rune切片

	//for _,c1 := range s2{
	//	fmt.Printf("%c\n", c1)
	//}
	s2[0] = '红'  // 单引号表示字符
	fmt.Println(string(s2))

	// 类型转换
	n1 := 10 // int
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("%T\n",f)

}
