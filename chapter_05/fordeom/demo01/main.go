package main

import "fmt"

func main(){
	// for第一种写法
	for i:=1; i <= 10; i++{
		fmt.Println("hello for01",i)	
	}

	// for第二种写法
	j := 1    // 循环变量初始化
	for j <= 10{  // 循环条件
		fmt.Println("hell for02",j)
		j ++   // 循环变量迭代
	}

	// 
	m := 1
	for { 
		fmt.Println("hell for03",m)
		m ++ 
		if m > 10 {
			break
		}
			
	}
		
}