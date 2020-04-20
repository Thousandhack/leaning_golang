package main

import "fmt"

func main() {
	// 白色  金色   紫色   青色
	// 判断输入的颜色是都存在数组中
	// 思路：
	// 1.定义一个数组 白色  金色   紫色   青色
	//2.从控制台接收一个名词，依次比较，如发现有，提示
	colors := [4]string{"白色", "金色", "紫色", "青色"}
	var color = ""
	// 顺序查找第一种方式
	fmt.Println("请输入一个颜色：")
	fmt.Scanln(&color)
	for i := 0; i < len(colors); i++ {
		if color == colors[i] {
			fmt.Println("找到了颜色：", color)
		} else if i == (len(colors) - 1) {
			fmt.Println("没有找到")
		}

	}
	// 顺序查找第二种方式
	index := -1
	for i := 0; i < len(colors); i++ {
		if color == colors[i] {
			index = i // 将找到的值对应的下标赋给index
			break
		}
	}
	if index != -1 {
		fmt.Println("找到了", color, index)
	} else{
		fmt.Println("没有找到")
	}
}
