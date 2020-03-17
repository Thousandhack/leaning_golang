package main

import "fmt"


func main(){
	// // 求一组数的平均值
	// num1 := 3.0
	// num2 := 5.0
	// num3 := 1.0
	// num4 := 3.4
	// num5 := 2.0
	// num6 := 50.0
	// total := num1 + num2 + num3 + num4 + num5 + num6
	// avg_num := fmt.Sprintf("%.2f",total / 6)
	// fmt.Printf("total=%v , avg_num=%v",total,avg_num)

	// 使用数组
	// 定义一个数组
	var nums [6]float64;
	// 给数组的每一个元素赋值
	nums[0] = 3.0
	nums[1] = 5.0
	nums[2] = 1.0
	nums[3] = 3.4
	nums[4] = 2.0
	nums[5] = 50.0

	// 遍历数字求平均值
	total := 0.0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	avg_num := fmt.Sprintf("%.2f",total / float64(len(nums)))
	fmt.Println(avg_num)
}