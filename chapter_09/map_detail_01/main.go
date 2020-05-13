package main

import "fmt"

func modify(map1 map[int]int){
	map1[0] = 100
}

func main(){
	// map是引用类型，遵守引用类型传递的机制
	// 修改后，会修改原来的map
	map1 := make(map[int]int)
	map1[0] = 90
	map1[1] = 80
	map1[2] = 70
	modify(map1)
	fmt.Println(map1)  // map在调用函数的时候被修改了
}