package main

import "fmt"

func main(){
	// 输入1-12 月份输出季节， 3,4,5 春季   6,7,8, 夏季  9,10,11 秋季 12,1,2 冬季
	var month int 
	fmt.Println("请输入月份:")
	fmt.Scanln(&month)
	switch int(month / 3) {
		case 1:
			fmt.Println("春季")
		case 2:
			fmt.Println("夏季")
		case 3:
			fmt.Println("秋季")
		default:
			fmt.Println("冬季")		
	}
}