package main

import "fmt"

// 声明/定义一个接口
type Usb interface {
	// 生命两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	Name string
}

// 让Phone实现use 接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")

}

func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")

}

// 让Camera 实现 Use 接口的方法
type Camera struct {
	Name string
}

func (computer Computer) Working(usb Usb){
	usb.Start()
	usb.Stop()
}

type Computer struct{

}

func (c Camera) Start() {
	fmt.Println("相机开始工作。。。")

}

func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")

}

func main(){
	// 定义一个use接口数组，可以存放Phone和Camera的结构体变量
	// 这里就体现出多态数组
	// 利用了接口多态的特点，原来一个数组只可以放一种类型，现在可以实现多种类型
	// 多态数组
	var useArr [3]Usb
	useArr[0] = Phone{"vivo"}
	useArr[1] = Phone{"mi"}
	useArr[2] = Camera{"hua"}
	// 遍历数组
	// Phone还有一个特有的方法call() ，请遍历Use数组，如果是Phone变量
	// 除了调用Usb 接口声明的方法外，还需要调用Phone 特有方法 call => 类型断言
	var computer Computer
	for _,v := range useArr{
		computer.Working(v)
	}

	fmt.Println(useArr)
}