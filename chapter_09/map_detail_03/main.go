package main

import "fmt"


// 定义一个学生结构体
type Stu struct{
	Name string
	Age  int
	Address string

}

func main(){
	// map的value的类型定义为struct类型
	// 
	students := make(map[string]Stu,10)
	stu1 := Stu{"hsz",20,"深圳"}
	stu2 := Stu{"lin",22,"杭州"}

	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)

	// 遍历学生信息
	for k,v := range students{
		fmt.Printf("k=%v,name=%v,age=%v,address=%v\n",k,v.Name,v.Age,v.Address)
	}

}