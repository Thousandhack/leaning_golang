package main

import "fmt"

func main()  {
	// 使用map 的key-value 特性 存放三个学生信息，包括姓名和年龄
	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string,2)
	studentMap["stu01"]["name"] = "zero"
	studentMap["stu01"]["age"] = "22"

	studentMap["stu02"] = make(map[string]string,2)
	studentMap["stu02"]["name"] = "one"
	studentMap["stu02"]["age"] = "23"

	studentMap["stu03"] = make(map[string]string,2)
	studentMap["stu03"]["name"] = "two"
	studentMap["stu03"]["age"] = "23"

	fmt.Println(studentMap)

}