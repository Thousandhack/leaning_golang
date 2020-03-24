package main

import "fmt"
import "math/rand"
import "time"


func main()  {
	// 一个数组的反转
	// 要求：随机生成5个数，并将其反转打印
	// rand.Intn()  生成（0，n）的随机数
	// 反转打印，交换的次数是 len / 2
	// 第一个与倒数第一  第二与倒数第二
	var intArr [5] int 
	// 为了每次生成的随机数不一样
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<len(intArr); i++ {
		intArr[i] = rand.Intn(100)  // 
	}
	fmt.Println(intArr)
	temp := 0
	for i:=0; i < len(intArr) / 2; i ++{
		temp = intArr[len(intArr) - 1 - i]
		intArr[len(intArr) -1 - i] = intArr[i]
		intArr[i] = temp
	}
	fmt.Println(intArr)
}