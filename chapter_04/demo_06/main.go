package main

import "fmt"

func main() {
	// 要求：可以从控制台接收用户信息，[姓名，年龄，薪水，是否通过考试]
	// 先声明需要的变量
	var name string
	var age byte
	var sal float32
	var isPass bool
	// 方式 2：fmt.Scanf,可以按指定的格式输入
	fmt.Println("请输入姓名，年龄，薪水，是否通过考试，使用空格隔开：")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf(" 姓名是：%v\n 年龄是：%v\n 薪水是：%v\n 是否通过考试：%v\n", name, age, sal, isPass)
}
