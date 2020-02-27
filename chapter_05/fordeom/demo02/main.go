package main

import "fmt"

func main(){
	// 字符串遍历方式1
	var str string = "hello,world"
	for i:=0; i < len(str); i++ {
		fmt.Printf("%c\n",str[i])
	}
	fmt.Println("===================================")
	// 字符串遍历方式2
	for index,val := range str {
		// index 表示的是下标   val 表示输出的内容
		fmt.Printf("index=%d,val=%c\n",index,val)
		
	}

	// 字符串遍历 中文不出错
	var str1 string = "hello,world,你好你好"
	str2 := []rune(str1)  // 就是str转成 []rune
	for i:=0; i < len(str2); i++ {
		fmt.Printf("%c\n",str2[i])
	}
}