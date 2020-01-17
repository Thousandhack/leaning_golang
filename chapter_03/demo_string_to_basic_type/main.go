package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "true"
	var b bool
	// 说明
	// 1.strconv.ParseBool(str) 函数会返回两个值(varlue bool,err error)
	// 2.因为我只想获取到value
	// _ 下划线就是忽略错误的意思
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n", b, b)

	// 字符串转成整数类型
	var str2 string = "123490"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n1 type %T b=%v\n", n1, n1)
	n2 = int(n1)
	fmt.Printf("n2 type %T b=%v\n", n2, n2)

	// 字符串转成float64类型
	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T b=%v\n", f1, f1)
	// 返回什么类型，就用什么类型接受，另外在转成的相同类型下的不同位
}
