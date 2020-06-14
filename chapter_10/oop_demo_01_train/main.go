package main

import "fmt"


/*
学生案例：
编写一个结构体，包含name gender age id score 字段 ，类型自己定义
结构体中声明一个say方法,返回string类型，方法返回信息包含学生的所有字段
在main方法中，创建stu结构体实例，并调用say方法，得到结果

*/

type Student struct{
	Name string
	Gender string
	Age int
	Id int
	Score float64
}

func (stu Student) say() string{
	infoStu := fmt.Sprintf("Student info name=%v gender=%v age=%v id=%v score=%v",stu.Name,stu.Gender,stu.Age,stu.Id,stu.Score)
	return infoStu
}


func main()  {
	var stu = Student{
		Name : "zero",
		Gender:"1",
		Age : 10,
		Id : 1,
		Score: 80.00,
	}
	info := stu.say()
	fmt.Println(info)
}