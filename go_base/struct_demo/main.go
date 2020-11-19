package main

import "fmt"

// 结构体是值类型
type person struct {
	name, gender string
}

func f(x person) {
	x.gender = "女"   // 修改的是副本的gender
}

func f2(x *person) {
	(*x).name = "zero"
	// 语法糖 下面一句和上面效果一样
	//*x.name = "zero"
}

func f_test(){
	var a int
	a = 100
	b := &a   // 将a 的指针地址赋值给b,所以b的值等于a的指针地址
	fmt.Printf("a type %T b type %T\n",a,b)
	fmt.Printf("%p\n",&a)  // a 的指针地址
	fmt.Println(*b)  // a的指针地址加*为a的值
	fmt.Printf("%p\n",b)   // b 的值
	fmt.Printf("%p\n",&b)   // b 的指针地址
}

type x struct {
	a int8
	b int8
	c int8
}

func structTestTwo(){
	m := x{
		a: int8(10),
		b: int8(20),
		c:int8(30),
	}
	fmt.Printf("%p\n",&(m.a))
	fmt.Printf("%p\n",&(m.b))
	fmt.Printf("%p\n",&(m.c))
}

func main() {

	f_test()

	var p person
	p.name = "hsz"
	p.gender = "男"
	fmt.Printf("%#v\n",p)
	fmt.Println(p)
	f(p)
	fmt.Println(p.gender) // 数据不变
	f2(&p)  // 传入的是指针
	fmt.Println(p.name)   // 因为引入的是指针地址，所以数据改变了

	// 结构体指针 1
	// p2是指针类型
	var p2 = new(person)
	fmt.Printf("%T\n",p2)
	p2.name = "two"
	fmt.Println()

	// 结构体指针2
	// 2.1 key-value 初始化
	var p3 = person{
		name:   "元帅",
		gender: "男",
	}

	fmt.Println(p3)
	// 2.2使用列表的形式初始化，值的顺序和结构体定义时字段的顺序一致
	var p4 = person{
		"four",
		"男",
	}
	fmt.Println(p4)
	// 内存地址的连续例子
	fmt.Println("下面的输出可以看出内存地址的连续：")
	structTestTwo()
	/*
	a type int b type *int
	0xc00006e090
	100
	0xc00006e090
	0xc000098018
	main.person{name:"hsz", gender:"男"}
	{hsz 男}
	男
	zero
	*main.person

	{元帅 男}
	{four 男}
	0xc00006e0e0
	0xc00006e0e1
	0xc00006e0e2

	*/

}
