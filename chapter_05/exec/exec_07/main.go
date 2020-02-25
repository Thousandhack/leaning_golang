package main

import "fmt"

func main() {
	// 嵌套分支
	// 题目：参加百米运动会，如果用时8秒以内进入决赛，否则提示淘汰。并且
	// 根据性别提示进入男子组或女子组。
	// 输入成绩和性别。
	// var time float64
	// fmt.Println("请输入秒数:")
	// fmt.Scanln(&time)
	// if time < 8 {
	// 	var sex string
	// 	fmt.Scanln(&sex)
	// 	if sex == "man" {
	// 		fmt.Println("男子跑8秒内,进入决赛")
	// 	} else if sex == "woman" {
	// 		fmt.Println("女子跑8秒内,进入决赛")
	// 	}

	// } else {
	// 	fmt.Println("没有进入决赛")
	// }

	// 出票系统：根据淡季旺季的月份和年龄，打印票价
	// 4-10 旺季：
	// 成人（16-60）：60
	// 儿童（<18）:半价
	// 老人（>60）：1/3

	// 淡季：
	// 成人： 40
	// 其他： 20
	var mouth byte
	var age byte
	var price float64 = 60.0
	fmt.Println("输入月份：")
	fmt.Scanln(&mouth)
	fmt.Println("输入年龄：")
	fmt.Scanln(&age)
	if mouth >= 4 && mouth <= 10 {
		if age > 60 {
			fmt.Printf("票价是%v", price/3)
		} else if age >= 18 {
			//
			fmt.Printf("票价是%v", price)
		} else {
			fmt.Printf("票价是%v", price)
		}
	} else {
		if age > 18 {
			fmt.Println("票价是40")
		} else {
			fmt.Println("票价是20")
		}
	}

}
