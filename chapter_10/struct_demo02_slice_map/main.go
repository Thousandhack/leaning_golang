package main

import "fmt"


// 如果结构体的字段为：指针，slice和map的默认值都是nil
// 如果需要使用这些字段，需要先make才能使用

type Person struct {
	Name string
	Age int
	Scores [5]float64   // 
	ptr *int  // 指针
	Slice []int  // 切片
	Map map[string]string   // map

}

func main(){
	// 定义结构体
	var person1 Person
	fmt.Println(person1)  // { 0 [0 0 0 0 0] <nil> [] map[]} 

	if person1.ptr == nil{
		fmt.Println("ok1")
	} 
	if person1.Slice == nil {
		fmt.Println("ok2")
	}
	if person1.Map == nil {
		fmt.Println("OK3")
	}

	// 使用slice先make
	person1.Slice = make([]int,10)
	person1.Slice[0] = 10
	fmt.Println(person1.Slice)

	// 使用map先make
	person1.Map = make(map[string] string)
	person1.Map["hh"] = "666"
	fmt.Println(person1.Map)
}