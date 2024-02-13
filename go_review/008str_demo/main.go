package main

import (
	"fmt"
	"strings"
)

// go语言字符串是双引号包裹的
// go语言单引号包裹的是字符

// 遍历字符串
func traversalString() {
	s := "hello深圳"
	// 第一种方式无法全部打印正确的字符
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func main() {
	// 字符串
	//s := "xia"

	// 单独的字母，汉字，符合表示一个字符
	//c1 := 'h'
	//c2 := '1'
	//c2 := '黄'

	// 字节： 1字节 = 8Bit （8个二进制位）
	// 1 个字符 'A' = 一个字节
	// 1 个utf-8编码的的汉字'黄' = 一般占3个字节
	path := "D:\\goCode\\goproject\\src\\learning_golang\\go_review\\008str_demo\\main.go"
	fmt.Printf(path)

	// 多行字符串
	s2 := `
		hhhh
		zzzz
		`
	fmt.Println(s2)

	// 字符串相关操作

	// 计算字符串长度
	fmt.Println(len("123"))
	// 字符串拼接
	d1 := "你好"
	d2 := "2024"
	d := d1 + d2
	fmt.Printf("%s\n", d)
	d3 := fmt.Sprintf("%s%s", d1, d2)
	fmt.Print(d3, "\n")
	// 分割
	ret := strings.Split("你好,hsz", ",")
	fmt.Println(ret)
	// [你好 hsz]
	fmt.Printf("%T\n", ret)
	// 包含 返回true false
	ss := "你好, hsz"
	// 是 rune 类型 也是 int32 类型
	for _, c := range ss {
		fmt.Printf("%c %T\n", c, c)
	}

	fmt.Println(ret)
	fmt.Println(strings.Contains(ss, "你好"))
	fmt.Printf("%T\n", ss)
	// 前缀 后缀
	fmt.Println(strings.HasPrefix(ss, "你好"))
	fmt.Println(strings.HasSuffix(ss, "你好"))
	// 获取第一次出现的字符的索引  注意：如果有中文的话，字符串中的一个字符是算3个索引位置
	s3 := "abcdb"
	fmt.Println(strings.Index(ss, "z"))
	fmt.Println(strings.Index(s3, "b"))
	// 最后出现一个索引的位置
	fmt.Println(strings.LastIndex(s3, "b"))
	// 将 ss 通过 ，分割在 通过 + 拼接
	// join 拼接操作 得到结果： 你好+ hsz
	fmt.Println(strings.Join(strings.Split(ss, ","), "+"))

	traversalString()
}
