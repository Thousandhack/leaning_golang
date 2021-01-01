package calc_demo

import "fmt"

func Add(x,y int) int{
	return x + y
}


func init(){
	fmt.Println("cale init函数自己执行了")
}