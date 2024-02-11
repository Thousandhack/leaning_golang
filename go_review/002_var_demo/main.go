package main

// 导入语句
import "fmt"

// 函数外只能防止标识符（变量/常量/函数/类型）的声明
// fmt.Printf("哈哈哈")  但是

// 程序的主入口函数
func main() {
	// 打印完指定的内容之后会在后面加一个换行符
	fmt.Println("哈哈哈")

	// 单独声明变量
	// var name string
	// var age int
	// var isOk bool
	var (
		name string
		age  int
		isOk bool
	)
	name = "hsz"
	age = 25
	isOk = false
	// go 语言中非全局变量声明了必须使用
	fmt.Print(name, age, isOk)
	fmt.Printf("%s,%d,%t", name, age, isOk)

}
