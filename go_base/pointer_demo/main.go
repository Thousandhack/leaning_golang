package main

import "fmt"

func newPointer(){
	var a = new(int)   // 初始化一个内存地址
	*a = 100    // 将内存地址赋值
	fmt.Println(*a)  // 根据内存地址取值
}

func main() {
	// 指针
	// Go语言中不存在指针操作，只需要记住两个符合
	// 1. &   取地址
	// 2. *   根据地址取值
	n := 18
	fmt.Println(&n) // 取n变量的内存地址
	p := &n
	fmt.Printf("%T\n", p)  // 表示整数的指针类型
	m := *p  // 根据地址取值
	fmt.Println(m)
	fmt.Printf("%T\n",m)
	/*
	打印：
		0xc000070090
		*int
		18
		int
	*/
	/*
	总结： 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
	*/

	newPointer()
}

/*
	make和new的区别
	1.make和new都是用来申请内存的
	2.new很少用，一般用来给基本数据类型申请内存：string、int,返回的是对应类型的指针（*string *int）
	3.make是用来给slice、map、chan申请内存的，make函数返回的是对应的这三个类型本身

*/