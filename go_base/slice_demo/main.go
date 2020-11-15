package main

import "fmt"

func makeSlice() {
	// 容量省略的时候就和长度一样
	s1 := make([]int, 5, 10) // 长度为 5  容量为 10
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}

func assignmentSlice() {
	// 切片的赋值
	s1 := []int{1, 3, 5}
	s2 := s1
	fmt.Println(s1, s2)
	// s1改变了s2也就改变了，都指向了同一个底层数组
	s1[2] = 7
	fmt.Println(s1, s2)
}

func forSlice() {
	s1 := []int{1, 2, 3}
	// 1.索引遍历
	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}

	// 2.for  range 循环
	for k, v := range s1 {
		fmt.Println(k, v)
	}
}

func appendSlice() {
	// 切片扩容，切片长度增加了底层数组就会换一个
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "广州") // 必须要接收
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "杭州", "珠海")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s2 := []string{"福州", "西安", "南京"}
	s1 = append(s1, s2...) // ... 表示把s2一个个元素拆开之后添加进去
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

}

func copySlice() {
	a1 := []int{1, 3, 5}
	a2 := a1 // 赋值数据一起改变
	var a3 = make([]int, 3, 3)
	copy(a3, a1) // copy 是不会一起改变的
	fmt.Println(a1, a2, a3)
	a1[1] = 100
	fmt.Println(a1, a2, a3)

	// 切片删除中间的元素元素
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1, cap(a1))
	// 数组
	x1 := [...]int{1, 3, 5}
	s1 := x1[:]
	fmt.Println(s1, len(s1), cap(s1))
	// 1.切片不保存具体的值
	// 2.切片对应一个底层数组
	// 3.底层数组都是占用一块连续的内存
	// 下面内存地址不变
	fmt.Printf("%p\n", s1)

	s1 = append(s1[:1], s1[2:]...)  // 修改了底层数组的数据
	fmt.Printf("%p\n", s1)
	fmt.Println(s1, len(s1), cap(s1))
	s1[0] = 100 // 改变了底层数组的值
	fmt.Println(x1) // 为什么是：[100 5 5]

}

func main() {
	// 切片的定义
	// 切片的本质
	// 切片不能比较
	//
	var s1 []int
	var s2 []string
	// nil 是为空的意思
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	// 1.初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"深圳", "南山", "宝安"}
	fmt.Println(s1)
	fmt.Println(s2)

	// 2.由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s3 := a1[2:7]
	// 切片指向一个底层的数组
	// 这种情况下的切片的最大容量是数组切片的第一个元素在数组的的元素位置到最后一个的容量
	fmt.Println("s3切片的数据：", s3)
	fmt.Println("s3切片的长度：", len(s3))
	fmt.Println("s3切片的容量：", cap(s3))
	// 切片再切片
	// 切片是引用类型，都指向了一个底层数组
	s4 := s3[1:3]
	a1[4] = 11 // 修改了底层数组
	fmt.Println("s4切片的数据：", s4)
	fmt.Println("s4切片的长度：", len(s4))
	fmt.Println("s3切片的容量：", cap(s4))

	// 3.make 函数创建切片
	// 可以自己确定长度和容量
	makeSlice()

	// 切片的赋值
	assignmentSlice()
	// 切片的遍历
	forSlice()
	// 切片的扩容
	appendSlice()
	/*
		切片的扩容策略
		首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
		否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
		否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
		如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
		需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如int和string类型的处理方式就不一样。
	*/
	// 拷贝切片
	copySlice()
}
