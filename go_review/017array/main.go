package main

import "fmt"

func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}

func main() {
	// 数组
	// 存放元素的容器
	// 必须指定存放的元素的类型和容量（长度）
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:%T , a2:%T,\n", a1, a2)
	fmt.Println(a1, a2)

	// 1.初始化方式1 默认元素都是零值
	a1 = [3]bool{false, true, false}
	fmt.Println(a1)

	//2. 初始化方式2  根据初始值自动推断的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println(a10)

	// 3.初始化方式3：根据索引来初始化
	// 有赋值的不是零值，其他的是
	a3 := [5]int{1: 3, 4: 2}
	a3[2] = 7
	fmt.Println(a3)

	// 数组的遍历
	citys := [...]string{"北京", "上海", "深圳"}
	// 1.根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆
	fmt.Println("数组的遍历")
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
		//fmt.Println()
	}

	// 数组的修改
	// 数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
	b := [3]int{10, 20, 30}
	modifyArray(b) //在modify中修改的是a的副本x
	fmt.Println(b) //[10 20 30]
	c := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(c) //在modify中修改的是b的副本x
	fmt.Println(c)  //[[1 1] [1 1] [1 1]]

}
