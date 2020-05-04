package main

import (
	"fmt"
	"math/rand"  
	"time" 
)

func RandInt() int{
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(101)  // 生成1-100的随机整数
	return x
}

func main(){

	// y := RandInt()
	// fmt.Println(y)
	var arr_demo [10] int
	for i := 0; i < len(arr_demo); i++ {
		y := RandInt()
		arr_demo[i] = y
		time.Sleep(1)
	}
	fmt.Println(arr_demo)
}