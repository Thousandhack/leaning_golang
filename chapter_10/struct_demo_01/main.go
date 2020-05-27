package main

import "fmt"

// // 定义一个结构体
type Cat struct{
	Name string 
	Age int
	Color string
}

func main() {
	// 1.变量
	var cat01Name string = "cat_hei"
	var cat01Age int = 3
	fmt.Println(cat01Name, cat01Age)
	var cat02Name string = "cat_bai"
	var cat02Age int = 4
	fmt.Println(cat02Name, cat02Age)
	// 2.数组
	var catNames [2]string = [...]string{"hei", "bai"}
	var catAges [2]int = [...]int{3, 4}
	fmt.Println(catNames, catAges)

	// 3.使用结构体
	var cat1 Cat   // 使用一个结构体
	cat1.Name = "bai"
	cat1.Age = 2
	cat1.Color = "white"
	// fmt.Println(cat1)
	fmt.Println("Name = ",cat1.Name)
	fmt.Println("Age = ",cat1.Age)
	fmt.Println("Color = ",cat1.Color)

}
