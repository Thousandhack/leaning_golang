package main

import "fmt"

func main()  {
	//二维数组的遍历
	var arr [2][3]int = [2][3]int {{1,2,22},{12,23,45}}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v\t",arr[i][j])
		}
		fmt.Println()
	}
}