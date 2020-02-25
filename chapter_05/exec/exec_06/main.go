package main

import "fmt"

func main() {
	// 嵌套分支
	// 题目：参加百米运动会，如果用时8秒以内进入决赛，否则提示淘汰。并且
	// 根据性别提示进入男子组或女子组。
	// 输入成绩和性别。
	var time float64
	fmt.Println("请输入秒数:")
	fmt.Scanln(&time)
	if time < 8 {
		var sex string
		fmt.Scanln(&sex)
		if sex == "man" {
			fmt.Println("男子跑8秒内,进入决赛")
		} else if sex == "woman" {
			fmt.Println("女子跑8秒内,进入决赛")
		}

	} else {
		fmt.Println("没有进入决赛")
	}
}
