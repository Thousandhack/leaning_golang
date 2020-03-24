package main

import "fmt"


func main(){
	// 请求一个数组的和和平均值。 for range
	var intArr[5] int = [...]int {1,-1,-5,90,11}
	sum := 0
	for _,val := range intArr{
		// 求和
		sum += val
	}
	fmt.Println(sum,float64(sum)/float64(len(intArr)))

}