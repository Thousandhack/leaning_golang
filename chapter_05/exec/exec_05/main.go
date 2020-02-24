package main

import "fmt"

func main() {
	// 男方条件：1.高180cm以上 2.富1000万   3.帅
	// 1）三个都满足，则：输出一定嫁给他
	// 2) 有一个为真： 嫁，比上不足，比下有余
	// 3) 都不满足：  不嫁
	var  hight float32
	var money int 
	var handsome bool
	fmt.Println("请输入身高(cm)：")
	fmt.Scanln(&hight)    
	fmt.Println("请输入财富(w)：")
	fmt.Scanln(&money)  
	fmt.Println("请输入是否帅：")
	fmt.Scanln(&handsome)  

	if hight > 180 && money > 1000 && handsome {
		fmt.Println("一定嫁给他")
	}else if hight > 180 || money > 1000 || handsome{
		fmt.Println("嫁，比上不足，比下有余")
	}else {
		fmt.Println("不嫁")
	} 

}
