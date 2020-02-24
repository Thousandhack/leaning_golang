package main

import "fmt"
import "math"

func main() {
	// 求ax2 + bx + c = 0 方程的根。a,b,c 分别为函数的参数
	// 如果：b2 - 4ac > 0 ,则有两个解
	// 如果：b2 - 4ac = 0 ,则有一个解
	// b2 - 4ac < 0 ,则无解
	// 提2： x1 = (-b+sqrt(b2-4ac)) / 2a
	//       x2 = (-b-sqrt(b2-4ac)) / 2a
	// 提示2： math.Sqrt(num)；需要求平方根 需要引入 math 包

	var a float64 = 2.0
	var b float64 = 4.0
	var c float64 = 2.0

	m := b*b - 4*a*c
	// 多分支判断
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		x2 := (-b - math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v   x2=%v", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		// x2 = x1
		fmt.Printf("x1=%v", x1)
	} else {
		fmt.Println("无解")
	}

}
