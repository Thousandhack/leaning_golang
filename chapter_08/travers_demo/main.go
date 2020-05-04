package main


import "fmt"

func main(){
	// 题目：定义二维数组，用于保存三个班，每个班五名同学成绩，
	// 并求出每个班级平均分，以及所有班级平均分。

	// 定义一个二维数组
	var scores [3][5] float64 
	// 循环输入成绩
	for i := 0 ; i < len(scores); i++ {
		for j:= 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第%d班的第%d个学生的成绩:",i+1,j+1)
			fmt.Scanln(&scores[i][j])
		}
	}

	fmt.Println(scores)
	// 遍历输出成绩的二维数组，统计平均分
	totalSum := 0.0
	for i := 0; i < len(scores); i++ {
		sum := 0.0  // 初始的每个班级总分值
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]
		}
		totalSum += sum
		fmt.Printf("第%v班级总分为%v,平均分为%v\n",i+1,sum,sum /float64(len(scores[i])))
	}
	fmt.Printf("所有班级的总分%v,所有班级的分均分%v\n",totalSum,totalSum / float64(15))
}