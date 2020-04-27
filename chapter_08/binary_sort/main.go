package main

import (
	"fmt"
)

// 二分查找
/*
二分查找的思路:比如要查找findval
1.arr是一个有序数组，并且是从小到大排序
2.先找到中间的下标mid = (leftindex + rightindex) / 2,然后让中间下标的值和findval进行比较
2.1 如果arr[mid] > findval,就应该leftindex -----(mid-1)
2.2如果arr[mid] < findval ，就应该mid + 1 ---- rightindex
2.3如果arr[mid] == findval，就找到
2.4上面的2.1 2.2 2.3的逻辑递归执行

3.想一下，怎么样的情况下，就说明找不到分析出退出递归的条件
*/
func BinarySort(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	// 判断leftIndex 是否大于rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	// 中间的下标
	mid := (leftIndex + rightIndex) / 2

	if (*arr)[mid] > findVal {
		// 说明要查找的数，应该在left leftIndex --- mid -1
		BinarySort(arr, leftIndex, mid-1, findVal)
	} else if (*arr)[mid] < findVal {
		// 说明要查找的数，应该在left mid + 1 ---- rightIndex
		BinarySort(arr, mid+1, rightIndex, findVal)
	} else {
		// 找到了
		fmt.Printf("下标为%v", mid)
	}

}

func main() {
	//
	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	BinarySort(&arr, 0, len(arr)-1, 1000)
}
