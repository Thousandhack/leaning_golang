package main

import "fmt"

func main() {
	// 2)对学生成绩大于60分，输出“合格”。低于60 输出不合格
	// 输入的成绩不大于100分
	var score float64
	fmt.Println("请输入成绩:")
	fmt.Scanln(&score)

	if score < 100 && score >= 0 {
		switch  int(score / 10) {
			case 10:
				fmt.Println("合格")
			case 9:
				fmt.Println("合格")
			case 8:
				fmt.Println("合格")
			case 7:
				fmt.Println("合格")
			case 6:
				fmt.Println("合格")
			case 5:
				fmt.Println("不合格")
			case 4:
				fmt.Println("不合格")
			case 3:
				fmt.Println("不合格")
			case 2:
				fmt.Println("不合格")
			case 1:
				fmt.Println("不合格")
			case 0:
				fmt.Println("不合格")
			default:
				fmt.Println("输入有误")
		}
	} else {
		fmt.Println("输入有误")
	}
}
