package main

import "fmt"

/*
Golang中的方法作用在指定的数据类型上的（即：和指定的数据类型绑定）,
因此自定义类型，都可以有方法，而不仅仅是struct,比如int,float32 等都可以有方法
*/

type integer int

func (i *integer) print(){ // 如果首字母方法大写，那么其他包也可以调用
	*i = *i + 1
	fmt.Println("i=",*i)
}

func main(){
	var i integer = 10
	i.print()
	fmt.Println(i)
}
