package main

import "fmt"

func sum(n1 int,args...int) int {
	sum := n1
	for i := 0; i < len(args); i ++ {
		sum += args[i]  // args[0] 表示第一个元素的值 其他以此类推
	} 

	return sum
}

func main(){
	// 编写一个函数sum ，可以求出1到多个int的和
	res := sum(5,6,7,8,1,2,3)
	fmt.Println(res)

}