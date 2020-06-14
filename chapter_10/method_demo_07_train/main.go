package main

import "fmt"

type MethodUtils struct {

}

func (mu MethodUtils) Print(m int ,n int){
	for i:=1; i <= m; i++{
		for j:=1; j <= n; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) Area(len float64,width float64) (float64){
	return len * width
}

func main(){
	// 编一个方法，提供m和n两个参数，方法打印m*n 的矩形
	// 方法中接收参数
	var mu MethodUtils
	mu.Print(5,6)
	res := mu.Area(4.0,4.0)
	fmt.Println("面积为：",res)
}
