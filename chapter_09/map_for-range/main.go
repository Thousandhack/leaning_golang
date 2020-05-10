package main

import "fmt"

func main() {

	// 使用for-range遍历map
	cities := make(map[string]string)
	cities["城市1"] = "北京"
	cities["城市2"] = "天津"
	cities["城市3"] = "上海"
	fmt.Println(cities)
	fmt.Printf("map有%v个\n", len(cities))

	for k, v := range cities {
		fmt.Printf("k=%v, v=%v\n", k, v)
	}

	// 双层map的遍历
	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 2)
	studentMap["stu01"]["name"] = "zero"
	studentMap["stu01"]["age"] = "22"

	studentMap["stu02"] = make(map[string]string, 2)
	studentMap["stu02"]["name"] = "one"
	studentMap["stu02"]["age"] = "23"

	studentMap["stu03"] = make(map[string]string, 2)
	studentMap["stu03"]["name"] = "two"
	studentMap["stu03"]["age"] = "23"

	fmt.Println(studentMap)

	for k1, v1 := range studentMap {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v,v2=%v\n", k2, v2)
		}
	}
}
