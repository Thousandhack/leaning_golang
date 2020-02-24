package main

import "fmt"

func main() {
	// 分析思路
	// 1.score 分数变量 int
	// 2.选择多分支流程控制
	// 3.成绩输入 fmt.Scanln
	var score int
	fmt.Println("请输入成绩：")
	fmt.Scanln(&score)

	// 多分支判断
	if score == 100 {
		fmt.Println("奖励一辆车")
	} else if score > 80 && score <= 99 {
		fmt.Println("奖励一手机")
	} else if score > 60 && score <= 80 {
		fmt.Println("奖励一手表")
	} else {
		fmt.Println("没有奖励")
	}

}
