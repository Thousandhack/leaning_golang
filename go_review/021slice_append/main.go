package main

import "fmt"

func main() {
	// 切片的append 为切片追加元素
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	//s1[3] = "广州" // 会索引越界
	//fmt.Println(s1)

	// 调用append函数必须用原来的切片变量接收返回值
	// append追加元素，原来的底层的数组放不下，Go底层就会把底层数组换一个
	s1 = append(s1, "广州") // 必须用变量接收返回值
	fmt.Println(s1)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "杭州", "成都")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	//
	ss := []string{"苏州", "潮州"}
	s1 = append(s1, ss...) // ... 表示拆开
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}
