package main

import "fmt"

func main() {
	// 如果一次性声明多个变量
	var n1, n2, n3 int
	n1 = 3
	n2 = 5
	n3 = 7
	fmt.Println("n1+n2+n3=", n1+n2+n3)
	// 法二：
	var m1, name, m3 = 100, "tom", 888
	fmt.Println("m1:", m1, "name:", name, "m3:", m3)
	// 法三：
	var (
		a = 1
		b = 2
		c = 3
	)
	fmt.Println("a=", a, "\n", "b=", b, "\n", "c=", c)
}
