package main

import "fmt"

func main() {
	// map
	var m1 map[string]int
	// 初始化
	m1 = make(map[string]int, 10)
	m1["hsz"] = 70
	m1["zero"] = 88
	fmt.Println(m1)
	fmt.Println(m1["hsz"])

	value, ok := m1["haha"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}
	// 27 8分钟
}
