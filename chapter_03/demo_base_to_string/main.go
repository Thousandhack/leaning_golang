package main

import "fmt"
import "strconv"

func main() {
	// 基本数据类型转String类型
	var num1 int = 99
	var num2 float64 = 23.234
	var bool_num bool = true
	var mychar byte = 'h'
	var str string
	// 第一种方式 fmt.Sprintf 方法
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q\n", str, str) // 输出类型和数值

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str) // 输出类型和数值

	str = fmt.Sprintf("%t", bool_num)
	fmt.Printf("str type %T str=%q\n", str, str) // 输出类型和数值

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type %T str=%q", str, str) // 输出类型和数值

	// 第二种方式 strconv 函数
	var num3 int = 99
	var num4 float64 = 23.234
	var bool_num2 bool = true
	// var mychar2 byte = 'g'
	var str2 string
	str2 = strconv.FormatInt(int64(num3),10)
	fmt.Printf("str2 type %T str=%q\n", str2, str2)
	// fmt.Printf("str type %T str=%q\n", str2, str2)
	// 说明： 'f' 格式 10，表示小数位表示十位 64 表示这个小数为float64

	str2 = strconv.FormatFloat(num4,'f',10,64) 
	fmt.Printf("str2 type %T str=%q\n", str2, str2)

	str2 = strconv.FormatBool(bool_num2) 
	fmt.Printf("str2 type %T str=%q\n", str2, str2)

}
