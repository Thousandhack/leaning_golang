package main

import "fmt"

func main() {
	// 回文判断
	// 字符串从右往左读和从左往右读是一样的，那么就是回文
	// 上海自来水来自海上
	// 山西运煤车煤运西山
	// 解题思路：
	// 把字符串中的字符拿出来放到一个[]rune切片中
	ss := "山西运煤车煤运西山"
	r := make([]rune, 0,len(ss))
	for _, c := range ss {
		r = append(r, c)
	}
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文字符串")
			return
		}
	}
	fmt.Println("是回文")

}
