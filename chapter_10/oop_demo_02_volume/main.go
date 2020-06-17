package main


import "fmt"

/*
	创建一个Box结构体，在其中什么三个字段表示一个立方体的长，宽和高，长宽高要从终端获取
	什么一个方法获取立方体的体积
	创建一个Box结构体变量，打印立方体的体积
*/

// 盒子结构体
type Box struct{
	length float64
	width  float64
	height float64
}

func (b Box) volume() float64{
	return b.length * b.width * b.height
}

func main(){
	// 实例化一个结构体

	var b Box
	b.length = 5.00
	b.width  = 4.00
	b.height = 3.00
	res  := b.volume()
	fmt.Println(res)

}
