package main

import "fmt"

func main(){
	// 统计3个班成绩情况。每个班5名同学，求出各个班的平均分和所有班级的平均分[学生的成绩自己输入]
	
	// var sum float64 = 0
	var class_num int = 3
	var stu_num int = 5
	total := 0.0
	for j := 1; j <= 3; j ++ {
		sum := 0.0
		for i:= 1; i <= 5; i++ {
			var score float64
			fmt.Printf("请输入第%d班第%d个学生的成绩:\n",j,i)
			fmt.Scanln(&score)
			sum += score
		}
		fmt.Printf("第%d班级的平均分是%v\n",j,sum / float64(stu_num))
		total += sum
	}
	fmt.Printf("各个班级的总成绩%v,各个班的平均分为%v",total,total/float64(stu_num)/float64(class_num))

}