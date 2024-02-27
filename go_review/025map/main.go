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
	// 约定俗称用ok 接收返回的布尔值
	value, ok := m1["haha"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}
	// map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 只遍历key
	for k := range m1 {
		fmt.Println(k)
	}
	// 只遍历v
	for _, v := range m1 {
		fmt.Println(v)
	}
	delete(m1, "hsz")
	fmt.Println(m1)
	// 删除没有的元素就没有反应
	delete(m1, "two")
}
