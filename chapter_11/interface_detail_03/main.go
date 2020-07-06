package main

import (
	"fmt"
)

// 接口
type Use interface {
	Say()
}

// 结构体
type Stu struct{

}

// 方法
func (this *Stu) Say(){
	fmt.Println("say ....")
}


func main(){
	var stu Stu = Stu{}
	var u Use = &stu
	u.Say()
}
