package main

import (
	"fmt"
	"os"
)

// 学生管理系统
// 有一个物件：
// 1.它保存了一些数据  ---> 结构体（其他语言的类）的字段
// 2.它有三个功能     --->  结构体的方法

var smr studentManager // 声明一个全局的变量（学生管理对象）



// 菜单函数
func showMenu() {
	fmt.Println("welcome sms !")
	fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.修改学生
		4.删除学生
		5.退出
		`)
}

func main() {
	smr = studentManager{
		allStudent: make(map[int64]student,100),
	}

	for {
		showMenu()
		fmt.Println("请输入你要做的序号:")
		// 2. 等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你输入的是：%d\n", choice)
		// 3. 执行对应的函数
		switch choice {
		case 1:
			smr.showStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.modifyStudent()
		case 4:
			smr.deleteStudent()
		case 5:
			os.Exit(1) // 退出
		default:
			fmt.Println("你的选择有问题")
		}
	}
}
