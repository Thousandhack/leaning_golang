package main

import (
	"fmt"
)

func test() bool {
	fmt.Println("test ok....")
	return true
}

func main() {
	// 演示关系运算符的使用
	var n1 int = 9
	var n2 int = 8
	fmt.Println(n1 == n2) // flase
	fmt.Println(n1 != n2) // true
	fmt.Println(n1 > n2)  // true
	fmt.Println(n1 >= n2) // true
	fmt.Println(n1 < n2)  // flase
	fmt.Println(n1 <= n2) // flase

	flag := n1 > n2
	fmt.Println("flag=", flag)

	// 逻辑运算符的使用  &&
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok1") // 执行
	}

	if age > 30 && age < 40 {
		fmt.Println("ok2") // 不执行
	}

	if age > 30 || age < 50 {
		fmt.Println("ok3") // 执行
	}

	if age > 30 || age < 40 {
		fmt.Println("ok4") // 执行
	}

	if !(age > 30 && age < 50) {
		fmt.Println("ok5") // 不执行
	}
	if !(age > 30 && age < 40) {
		fmt.Println("ok6") // 执行
	}

	var i int = 10
	fmt.Println("test01=========")
	if i > 9 && test() {
		fmt.Println("0k...")
	} // 这个先打印 test ok .... 再打印 ok...
	fmt.Println("test02==========")
	if i < 9 && test() {
		fmt.Println("0k...")
	} // 不打印   短路与现象

	fmt.Println("test03========")
	if i > 9 || test() {
		fmt.Println("0k...")
	} // 只打印ok   短路或现象
	fmt.Println("test04========")
	if i < 9 && test() {
		fmt.Println("0k...")
	} // 这个先打印 test ok .... 再打印 ok...

}
