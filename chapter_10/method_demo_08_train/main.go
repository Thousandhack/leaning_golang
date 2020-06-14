package main

import "fmt"

type MethodUtils struct {

}

func (mu MethodUtils) IsDouble(num int) (bool){
	// 判断是奇数还是偶数的方法
	if num % 2 == 0{
		return true
	}
	return false
}

func (mu MethodUtils) Print(m int,n int ,key string) {
	for i:=1; i<=n; i++ {
		for j:=1; j<=m; j++{
			fmt.Printf(key)
		}
		fmt.Println()
	}
}



func main(){
	// 编写一个方法，判断是奇数还是偶数
	var mu MethodUtils
	res :=mu.IsDouble(8)
	fmt.Println(res)

	// 实现打印几行几列，并可以自由指定想要打印的符号
	mu.Print(3,5,"(-_-)")
}


