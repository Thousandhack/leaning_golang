package main

import "fmt"
import "unsafe"

func main() {
	// 测是int8范围
	var i int8 = -128
	// var j int8 = -129  // 这个超出范围会报错
	// .\main.go:9:15: constant -129 overflows int8
	fmt.Println("i:", i)
	// fmt.Println("j:",j)

	// 测试uint8的范围
	var a uint8 = 255
	fmt.Println("a=", a)
	// .\main.go:17:16: constant 256 overflows uint8
	// var b uint8 = 256
	// fmt.Println("b=",b)

	// 测试数据类型 注意这个打印是：Printf
	v1 := "中国你好"
	v2 := 20
	var v3 byte = 65
	fmt.Printf("v1的数据类型为：%T\n", v1)
	fmt.Printf("v2的数据类型为：%T\n", v2)
	fmt.Printf("v2的数据类型为：%T\n", v3)

	// 查看某个变量的字节大小和数据类型
	var v int8 = 10
	fmt.Printf("v的类型 %T v占用的字节数是是%d", v, unsafe.Sizeof(v))
}
