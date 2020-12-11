package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {
	// 从字符串中解析出对应的数据
	// 将字符串中的数字转成int64
	int64Str := "10000"
	int64Value,err := strconv.ParseInt(int64Str, 10, 64) // 10进制，64位,但是通过源码可以知道，最后返回就是int64
	if err != nil{
		fmt.Println("ParseInt fail err:",err)
		return
	}
	fmt.Printf("%#v %T\n", int64Value,int64Value)

	// 将字符串转成 int 类型
	intValue,_ := strconv.Atoi(int64Str)
	fmt.Printf("%#v %T\n", intValue,intValue)

	i := int32(97)
	ret := fmt.Sprintf("%d", i)
	fmt.Printf("%#v\n", ret)

	// 将int型数字转换成字符串
	ret3 := strconv.Itoa(int(i))
	fmt.Printf("%#v %T\n", ret3,ret3)

	// 从字符串中解析出
	boolStr := "true"
	boolValue,_ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue,boolValue)

	// 将字符串中解析出浮点数
	floatStr := "12.345"
	floatValue,_ := strconv.ParseFloat(floatStr,64)
	fmt.Printf("%#v %T\n", floatValue,floatValue)

}
