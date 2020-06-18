package main

import (
	"fmt"
	"gotest01/chapter_10/factory_model/model"
	"gotest01/chapter_10/factory_model/teacher"
)

func main(){
	// 创建要给Student示例  正常的情况 结构体大写
	var stu = model.Student{
		Name : "hsz",
		Score:88.5,
	}
	fmt.Println(stu.Name,stu.Score)

	// 当teacher 结构体为首字母小写 ，工厂模式如下
	var t = teacher.NewTeacher("lin",28)    // 放回的是指针
	fmt.Println(*t)
	fmt.Println(t.Name,t.GetAge())  // 变量值小写的情况

}
