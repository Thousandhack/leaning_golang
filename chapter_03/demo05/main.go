package main

import "fmt"

func main() {
	// 该区域的数据值可以在同一类型范围内不断变化
	var i int = 10
    var j int = 20

    var o = i + j  // 做加法运算
    fmt.Println("o=",o)
    var m = "zero"
    var n = "one"
    var p = m + n
    fmt.Println("p=",p)  // 做字符串拼接
}