package main

import "fmt"

// 声明/定义一个接口
type Use interface{
	// 生命两个没有实现的方法
	Start()
	Stop()
}


//type Use_02 interface{
//	// 生命两个没有实现的方法
//	Start()
//	Stop()
//}



type Phone struct{

}


// 让Phone实现use 接口的方法
func (p Phone) Start(){
	fmt.Println("手机开始工作。。。")

}

func (p Phone) Stop(){
	fmt.Println("手机停止工作。。。")

}

// 让Camera 实现 Use 接口的方法
type Camera struct{

}


func (c Camera) Start(){
	fmt.Println("相机开始工作。。。")

}

func (c Camera) Stop(){
	fmt.Println("相机停止工作。。。")

}


// 计算机
type Computer struct{

}

// 编写一个方法 working
// 接收一个接口类型的数据
// 所谓实现use接口，就是实现了 use接口声明的所有方法
func (cm Computer) Working(use Use){
	// 通过接口调用方法
	use.Start()
	use.Stop()
}

func main(){

	// 测试
	// 创建结构体变量
	cm := Computer{}
	phone := Phone{}
	//camera := Camera{}

	//
	cm.Working(phone)
	//cm.Working(camera)
}