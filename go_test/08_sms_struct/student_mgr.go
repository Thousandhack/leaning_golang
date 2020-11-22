package main

import "fmt"

// 学生管理系统
// 有一个物件：
// 1.它保存了一些数据  ---> 结构体（其他语言的类）的字段
// 2.它有三个功能     --->  结构体的方法
type student struct{
	id int64
	name string
}

//var (
//	allStudent map[string]student
//)

// 学生管理者  对象
type studentManager struct {
	allStudent map[int64]student
}

// 查看学生
func (s studentManager) showStudent(){
	for _,stu := range s.allStudent{
		// 从s.allStudent这个map中所有学生进行遍历
		fmt.Printf("学号：%d 姓名：%s\n",stu.id,stu.name)
	}
}
// 增加学生
func (s studentManager) addStudent(){
	// 1.根据用户输入的内容创建一个新的学生
	var (
		StuID int64
		StuName string
	)
	// 获取输入的数据,根据用户输入创建结构体对象
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&StuID)
	fmt.Print("请输入学生名字:")
	fmt.Scanln(&StuName)
	newStu := student{id:StuID,name:StuName}
	// 2. 把新的学生放到s.allStudent这个map中
	s.allStudent[newStu.id] = newStu
}

// 修改学生
func (s studentManager) modifyStudent(){
	var (
		StuID int64
		ModifyName string
	)
	// 获取用户输入的学号
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&StuID)
	// 展示学号对应的学生信息，如果没有提示查无此人
	stuObj,ok :=s.allStudent[StuID]
	if !ok{
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("你要修改的学生信息如下：学号：%d 姓名：%s\n",stuObj.id,stuObj.name)
	fmt.Print("请输入修改学生名字:")
	fmt.Scanln(&ModifyName)
	// 更新学生的姓名
	//stuObj.name = ModifyName
	//s.allStudent[stuObj.id] = stuObj

	newStu := student{id:StuID,name:ModifyName}
	s.allStudent[newStu.id] = newStu

}
// 删除学生
func (s studentManager) deleteStudent(){
	var StuID int64
	// 获取要删除用户输入的学号
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&StuID)
	// 展示学号对应的学生信息，如果没有提示查无此人
	stuObj,ok :=s.allStudent[StuID]
	if !ok{
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("删除的学生信息如下：学号：%d 姓名：%s\n",stuObj.id,stuObj.name)
	delete(s.allStudent, StuID)
	fmt.Println("删除成功")
}
