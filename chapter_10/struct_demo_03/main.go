package main

import "fmt"

type Cat struct{
	Name string
	Age int
}


func main(){

	//  结构体里面的属性之间不会相互影响
	var cat1 Cat
	cat1.Name = "bai"
	cat1.Age = 2
	cat2 := cat1
	cat2.Age = 3
	fmt.Println(cat1)
	fmt.Println(cat2)
}