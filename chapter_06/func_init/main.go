package main

import "fmt"

var age = test()


func test() int {
	fmt.Println("run test")
	return 90
}

func init() {
	fmt.Println("run init")
}

func main(){
	fmt.Println("run main")
	fmt.Println("age=",age)
}

// 得出结论：先执行变量定义，再执行test-->init -- > main
// 打印结果：
// run test
// run init
// run main
// age= 90