package main

import "fmt"

// 编写一个学生考试的系统demo
// 匿名结构体
type Student struct{
	Name string
	Age int
	Score int
}

// 打印考试的信息
func (stu *Student) ShowInfo(){
	fmt.Println(stu.Name,stu.Age,stu.Score)
}

// 小学生
type Pupil struct{
	Student
}

// 大学生
type Graduate struct{
	Student
}


func (stu *Student) SetScore(score int){
	stu.Score = score
}

// 特有的方法
func (p *Pupil) p_testing(){
	fmt.Println("小学生正在考试。。。")
}

// 特有的方法
func (g *Graduate) g_testing(){
	fmt.Println("大学生正在考试。。。")
}

func main(){
	var pupil = &Pupil{}

	pupil.Student.Name = "test_one"
	pupil.Student.Age = 8
	pupil.p_testing()
	pupil.Student.SetScore(85)
	pupil.Student.ShowInfo()

	// 对结构体嵌入匿名结构体后，使用方法发生变化

	graduate := &Graduate{}
	graduate.Student.Name = "test_two"
	graduate.Student.Age = 22
	graduate.g_testing()
	graduate.Student.SetScore(90)
	graduate.Student.ShowInfo()


}
