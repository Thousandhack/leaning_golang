package main

import "fmt"

func main() {
	// 要求：可以输入人的年龄，如果该同志的年龄大于18岁，则输出，“你年龄大于18，要对自己的行为负责”。
	// 否则，输出"你的年龄不大，这次放过你"
	var age int
	fmt.Println("请输入你的年龄:")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("你年龄大于18，要对自己的行为负责")
	} else { // 注意这边不换行
		fmt.Println("你的年龄不大，这次放过你")
	}
}
