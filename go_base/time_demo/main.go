package main

import (
	"fmt"
	"time"
)

// 时间相关例子一
func timeDemoOne() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 时间戳的例子
func timestampDemo() {
	now := time.Now()
	timestampOne := now.Unix()     // 时间戳（秒）
	timestampTwo := now.UnixNano() // 时间戳 （纳秒）
	fmt.Println(timestampOne)
	fmt.Println(timestampTwo)
}

// timestamp to time 将一个时间戳转为时间
func timestampToTime(timestamp int64) {

	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

}

// time operate
func timeOperateDemo(){
	now := time.Now()
	later := now.Add(time.Hour) // 一小时后
	fmt.Println(later)
	fmt.Println(now.Add(time.Minute*10)) // 10分钟后
	// 定时器  可以做一个轮询的东西
	timer := time.Tick(time.Second)
	for t := range timer{
		fmt.Println(t)
	}
	//now.Sub()
}

func main() {
	timeDemoOne()
	timestampDemo()
	timestampToTime(1606788187)
	// 时间操作
	timeOperateDemo()
	//  https://www.liwenzhou.com/posts/Go/go_time/    78  6分钟
}
