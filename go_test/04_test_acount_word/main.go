package main

import (
	"fmt"
	"strings"
)

func main(){
	// 写一个程序，统计一个字符串中每个单词出现的次数。比如： "how do you do" 中how =1 do=2 you=1
	// 1. 把字符串安装切割得到切片
	// 2. 遍历切片存储到一个map
	// 3. 累加出现的次数
	s1 := "how do you do"
	s2 := strings.Split(s1," ")
	fmt.Println(s2)  // 数组
	m1 := make(map[string]int,10)
	for _,w := range s2{
		//fmt.Println(w)
		// 如果原来map中不存在,那么次数为1
		// 如果map中存在w这个key,那么次数加1
		if _,ok := m1[w];!ok{
			m1[w] = 1
		} else {
			m1[w] ++
		}
	}
	fmt.Println(m1)

}

