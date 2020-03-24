package main

import "fmt"

func  main()  {
	// 请求出一个数组最大值，并得到对应的下标

	// 思路 假定第一个元素就是最大值，下标是0
	// 然后从第二个元素开始循环比较，如果发现有更大，则交换
	var intArr[5] int = [...]int {1,-1,-99,90,11}
	maxVal := intArr[0]
	maxValindex := 0

	for i:=1; i<len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValindex = i
		}
		
	}
	fmt.Printf("maxVal=%v maxValindex=%v",maxVal,maxValindex)
}