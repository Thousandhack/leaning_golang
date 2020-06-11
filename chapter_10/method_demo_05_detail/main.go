package main

import "fmt"

/*
Golang中的方法作用在指定的数据类型上的（即：和指定的数据类型绑定）,
因此自定义类型，都可以有方法，而不仅仅是struct,比如int,float32 等都可以有方法
*/

type integer int

func (i *integer) print() { // 如果首字母方法大写，那么其他包也可以调用
	*i = *i + 1
	fmt.Println("i=", *i)

}

type Student struct {
	Name string
	Age  int
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name is %v Age is %v", stu.Name, stu.Age)
	return str
}

func main() {
	var i integer = 10
	i.print()
	fmt.Println(i)

	// 赋值一个stu变量
	stu := Student{
		Name: "ZEOR",
		Age:  25,
	}
	// 用指针
	/*
		如果一个类实现了string()这个方法,那么fmt.Print()默认调用这个变量的String进行输出 。
	*/
	fmt.Println(&stu)

}
