package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type RectPointer struct {
	leftUp, rightDown *Point
}

func main() {
	//
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	// r1 有4个int ，在内存中是连续分布的
	fmt.Println("r1.leftUp.x地址：", &r1.leftUp.x)
	fmt.Println("r1.leftUp.y地址：", &r1.leftUp.y)
	fmt.Println("r1.rightDown.x地址：", &r1.rightDown.x)
	fmt.Println("r1.rightDown.x地址：", &r1.rightDown.y)
	// // 八进制 一个地址八个字节
	/*
		r1.leftUp.x地址： 0xc000054120
		r1.leftUp.y地址： 0xc000054128
		r1.rightDown.x地址： 0xc000054130
		r1.rightDown.x地址： 0xc000054138
	*/
	fmt.Println("结构体下面的值为地址的情况")
	r2 := RectPointer{&Point{1, 2},&Point{3, 4}}
	fmt.Println("r2.leftUp.x地址：", &r2.leftUp.x)
	fmt.Println("r2.leftUp.y地址：", &r2.leftUp.y)
	fmt.Println("r2.rightDown.x地址：", &r2.rightDown.x)
	fmt.Println("r2.rightDown.x地址：", &r2.rightDown.y)
	/*
		结构体下面的值为地址的情况
		r2.leftUp.x地址： 0xc0000100a0
		r2.leftUp.y地址： 0xc0000100a8
		r2.rightDown.x地址： 0xc0000100b0
		r2.rightDown.x地址： 0xc0000100b8
		3
		0xc0000341d0
		0xc0000341d8
	*/
	fmt.Println(r2.rightDown.x)  //  x值
	// 
	fmt.Println(&r2.leftUp)
	fmt.Println(&r2.rightDown)
	
}
