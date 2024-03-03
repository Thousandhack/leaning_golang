package main

import "fmt"

// slice 和 map组合
func sliceMap() {
	var mapSlice = make([]map[string]string, 10)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "hsz"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "大沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

func mapSlice() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

func main() {
	sliceMap() // 元素为map类型的切片
	mapSlice() //值为切片类型的map
}
