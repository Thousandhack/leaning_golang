package main

import "fmt"


type Person struct{
	Name string
}

func (p Person) name(){
	p.Name = "test"
	fmt.Println("p.Name is ",p.Name)
}

func (p Person) add(){
	res := 0
	for i := 1; i<= 1000; i++{
		res += i
	}
	fmt.Println(p.Name,"add sum =",res)
}

func (p Person) two_sum(n1 int ,n2 int) int {
	return n1 + n2
}

func main(){
	//
	var p Person
	p.Name = "zero"
	p.name()   // name 方法无返回值
	p.add()    // add 方法无返回值
	res := p.two_sum(5,10) // 有返回值
	fmt.Println("res=",res)

}