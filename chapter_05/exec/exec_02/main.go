package main

import "fmt"

func main() {
	// 判断一个年份是否是闰年，闰年的条件符合下面二者之一
	//(1)年份能被4整除 ，但不能被100整除；（2）能被400整除
	var year int = 2019
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println(year, "年是闰年")
	} else {
		fmt.Println(year, "年不是闰年")
	}
}
