package main

import (
	"fmt"
	"time"
)

func main() {
	// 日期和时间相关函数与方法的使用
	// 1.获取当前时间
	now := time.Now()
	fmt.Printf("当前时间为:now=%v\n时间类型为：time_type=%T\n",now,now)
	// 打印结果：
	// 当前时间为:now=2020-03-15 17:59:06.3353254 +0800 CST m=+0.002996901
	// 时间类型为：time_type=time.Time

	// 通过now获取到年月日，时分秒
	fmt.Printf("年=%v\n",now.Year())
	fmt.Printf("月=%v  %v\n",now.Month(),int(now.Month())) // 英文的  数字加int强转
	fmt.Printf("日=%v\n",now.Day())
	fmt.Printf("时=%v\n",now.Hour())
	fmt.Printf("分=%v\n",now.Minute())
	fmt.Printf("秒=%v\n",now.Second())
	fmt.Printf("周=%v\n",now.Weekday())

}
