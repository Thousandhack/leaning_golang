package main

import (
	"fmt"
	"time"
)

func main() {
	// 格式化日期时间
	now := time.Now()
	fmt.Printf("当前年月日 时分秒 %02d-%02d-%02d %02d：%02d：%02d\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf("当前年月日 时分秒 %02d-%02d-%02d %02d：%02d：%02d\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Printf("%v %T\n", dateStr, dateStr)

	// "2006/01/02 15:04:05" 这个字符串的各个数字是固定的，必须是这样写
	// "2006/01/02 15:04:05" 这个字符串各个数字可以自由的组合，这样可以按程序需求来返回时间和日期
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	// 只输出年月日
	fmt.Println(now.Format("2006/01/02"))
	// 只输出时分秒
	fmt.Println(now.Format("15:04:05"))

	// fmt.Println(time.Unix)

	i := 0
	for {
		i ++
		fmt.Println(i)
		time.Sleep(time.Second)
		if i >= 10 {
			break
		}
	}


}
