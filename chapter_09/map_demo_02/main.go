package main

import "fmt"

func main() {
	// 方式一： 先声明再make
	var countries map[string]string
	// 在使用map前需要make 作用是给map分配数据空间
	countries = make(map[string]string, 10) // 10代表10个空间
	countries["国家1"] = "珠海"
	fmt.Println(countries)

	// 方式二：
	cities := make(map[string]string)
	cities["城市1"] = "北京"
	cities["城市2"] = "天津"
	cities["城市3"] = "上海"
	fmt.Println(cities)

	// 方式三：)
	hes := map[string]string{
		"he1": "one",
		"he2": "two",
	}
	fmt.Print(hes)
}
