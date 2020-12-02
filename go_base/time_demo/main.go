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
	// 定时器的本质上是一个通道（channel）。
	timer := time.Tick(time.Second)
	for t := range timer{
		fmt.Println(t)
	}
	//now.Sub()
}

func formatDemo() {
	now := time.Now()
	/*
	时间类型有一个自带的方法Format进行格式化，
	需要注意的是Go语言中格式化时间模板不是常见的
	Y-m-d H:M:S而是使用Go的诞生时间2006年1月2号15点04分
	（记忆口诀为2006 1 2 3 4）。也许这就是技术人员的浪漫吧。
	*/
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
	// 按照对应的格式解析字符串类型的时间
	timeObj,err := time.Parse("2006-01-02","2020-08-03")
	if err != nil{
		fmt.Printf("parse time failed,err:%v\n",err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
}

func main() {
	timeDemoOne()
	timestampDemo()
	timestampToTime(1606788187)
	// 时间操作
	//timeOperateDemo()
	formatDemo()
}
