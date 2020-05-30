package main

import "fmt"

type Person struct{
	Name string
	Age int
}

func main(){
	// 方式1：var p1 Person
	var p1 Person
	p1.Name = "one"
	p1.Age = 19
	fmt.Println(p1)

	// 方式2：p2 := Person{"two",20}
	p2 := Person{"two",20}
	fmt.Println(p2)

	// 方式3：
	var p3 *Person = new(Person)
	// 因为p3是个指针，因此标准的写法如下：
	(*p3).Name = "Three"
	(*p3).Age = 21
	fmt.Println(p3)
	// 可简化为:因为创建者底层对上面指针做了处理
	// p3.Name = "Three"
	// p3.Age = 21

	// 方式4：var p4 *Person = &Person()\
	// p4为指针
	var p4 *Person = &Person{}  // 也可以直接在{}进行赋值
	(*p4).Name = "four"
	(*p4).Age = 22
	fmt.Println(p4)
	fmt.Println(p4.Name)
	fmt.Println(p4.Age)
}