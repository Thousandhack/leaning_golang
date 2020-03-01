package main

import "fmt"

func main() {
	// 实现登录验证，有三次机会，如果用户名“张无忌”，密码“888” 提示登录成功，否则提示还有几次机会
	var name string
	var pwd string
	var login_chance int = 3
	for i := 1; i <= 3; i++ {
		fmt.Println("请输入用户名:")
		fmt.Scanln(&name)
		fmt.Println("请输入密码:")
		fmt.Scanln(&pwd)

		if name == "张无忌" && pwd == "888" {
			fmt.Println("恭喜你登录成功")
			break
		} else {
			login_chance --
			fmt.Printf("你还有%v次登录机会",login_chance)
		}

		if login_chance == 0 {
			fmt.Println("机会使用完了")
		}

	}
}
