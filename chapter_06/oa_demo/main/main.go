package main

import (
	"fmt"
	"gotest01/chapter_06/oa_demo/utils"   // 默认从src开始 
	// 注意这个utils是文件夹的名字
)

func main() {
	// 
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '-'
	result := utils.Cal(n1,n2,operator) // 这个utils是utils.go下package 后面的名字
	fmt.Println("result=",result)

}