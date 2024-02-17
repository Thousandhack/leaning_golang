package main

import "fmt"

func main() {
	//
	a := make([]int, 2, 10)
	fmt.Println(a)      //[0 0]
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10

	// 切片不能比较
	// 切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。 切片唯一合法的比较操作是和nil比较。

	// 切片的赋值
	s3 := []int{1, 3, 5}
	s4 := s3
	fmt.Println(s3, s4)
	s3[0] = 1000
	fmt.Println(s3, s4)

	// 切片的遍历
	// 1.索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	// 2. for range 循环
	for i, v := range s3 {
		fmt.Println(i, v)
	}

}
