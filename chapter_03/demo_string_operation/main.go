package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串相关操作
	// 打印一个路径
	var path string = "D:\\goCode\\goproject\\src\\gotest01\\chapter_03\\demo_string"
	fmt.Println("path:", path)
	fmt.Println("有多长:", len(path))
	// 通过// 将字符串变量path进行分割形成一个数组
	// 分割
	res := strings.Split(path, "\\")
	fmt.Println(res)
	fmt.Printf("%T\n", res)
	// 是否包含
	fmt.Println(strings.Contains(path,"go"))
	// 前缀和后缀判断
	fmt.Println(strings.HasPrefix(path,"D"))
	fmt.Println(strings.HasSuffix(path,"D"))

	path_son := "D:\\goCode\\goproject"
	// 返回字母在字符串中出现第一次对应的索引
	fmt.Println(strings.Index(path_son,"g"))
	// 返回字母在字符串中出现最后一次对应的索引
	fmt.Println(strings.LastIndex(path_son,"g"))

	// 拼接 把上面的res 数组里面的字符串进行 + 号拼接
	fmt.Println(strings.Join(res,"+"))

}
