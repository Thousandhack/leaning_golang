package main

import "fmt"

func main() {
	// map 的声明
	var a map[string]string
	// 在使用map前需要make 作用是给map分配数据空间
	a = make(map[string]string, 10) // 10代表10个空间
	a["age"] = "25"
	fmt.Println(a)

	fmt.Println(a["age"])
	// key不能重复 value会被重新赋值
	a["age"] = "24"
	fmt.Println(a)

	// key不重复的时候 value是可以重复的
	a["age_two"] = "24"
	fmt.Println(a)
}
