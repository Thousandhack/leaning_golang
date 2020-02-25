package main

import "fmt"

func main() {
	// 案例
	// 请编写一个程序，该程序可以接收一个字符，比如：a,b,c,d,e,f,g
	// a 表示星期一，b 表示星期二... 根据用户的输入显示相依据的信息

	// 要求使用switch 语句完成
	// 祝福一般用byte来接收
	var key byte
	fmt.Println("请输入字符:")
	fmt.Scanf("%c", &key)
	fmt.Println(key)
	switch key {

	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	case 'd':
		fmt.Println("周四")
	case 'e':
		fmt.Println("周五")
	case 'f':
		fmt.Println("周六")
	case 'g':
		fmt.Println("周七")
	default:
		fmt.Println("输入有误")
	}

	var age int = 10

	switch {
	case age == 10:
		fmt.Println("等于10")
	case age == 20:
		fmt.Println("等于20")
	case age == 30:
		fmt.Println("等于30")
	default:
		fmt.Println("没有匹配到")
	}

	// 穿透fallsthrough
	var num int = 10
	switch num {
	case 10:
		fmt.Println("10")
		fallthrough
	case 20:
		fmt.Println("20")
	case 30:
		fmt.Println("30")

	}
}
