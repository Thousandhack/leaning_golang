package main

import "fmt"

func main() {
	// GO语言变量使用注意事项：
	// 第一种：指定变量类型，声明后若不赋值，使用默认值
	var i int
	fmt.Println("i:", i)
	// 第二种：根据值，自行判断类型（类型推导）
	var num = 10.11
	fmt.Println("num:", num)
	// 第三种：省略var,注意： = 左侧的变量不应该是已经声明过的，否则会导致编译错误
	// var name string
	// name = "tom"
	name := "tom" // 与上面两行代码等价
	fmt.Println("name:", name)

}
