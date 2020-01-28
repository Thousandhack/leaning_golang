package main

import (
	"fmt"
	// 为了使用model.go文件的变量或者函数，我们需要先引入该包
	"gotest01/chapter_03/demo_standdard/model"
)

func main() {
	// 使用model.go文件的变量的HeroName可以导入过来
	fmt.Println(model.HeroName)
}
