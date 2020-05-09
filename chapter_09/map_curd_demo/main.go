package main

import "fmt"

func main(){
	classesMap := make(map[string]string)

	// 增加key-value数据
	classesMap["classOne"] = "班级1"
	classesMap["classTwo"] = "班级2"
	classesMap["classThree"] = "班级3"
	fmt.Println(classesMap)

	// 改map中的数据
	classesMap["classTwo"] = "班级2-2"
	fmt.Println(classesMap)

	// 删除map中的部分数据
	delete(classesMap,"classOne")  // 删除了classOne的数据
	fmt.Println(classesMap)
	// 删除全部数据两种方法：
	// 1.遍历所有数据，逐一删除
	// 2.map := make(...) 重新定义，让原来称为垃圾，让gc回收
	classesMap = make(map[string]string,0)
	fmt.Println(classesMap)

	// 查classTwo的数据
	classesMap["classThree"] = "班级3"
	fmt.Println(classesMap["classThree"])
	_, findRes :=  classesMap["classThree"]
	if findRes{
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
	// fmt.Println(classesMap["classThree"])

}