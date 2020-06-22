package main

// 引包从src开始
import (
	"fmt"
	"gotest01/chapter_10/oop_demo_06_encapsulation/model"
)

func main(){

	//
	p := model.NewPerson("test_one")
	fmt.Println(p)
	fmt.Println(p.Name)
	p.SetAge(18)    // Set的方法是设置结构体里面的值
	p.SetSalary(8800)
	fmt.Println(p)       // Get的方法是获取结构体的字段的值
	fmt.Println(p.GetAge(),p.GetSalary())
}
