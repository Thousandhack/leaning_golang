package main

import "fmt"

func main() {
	// 某人用100 000 元，每经过一次路口，需要交费规则如下
	// 1.当现金>50000,每次交5%
	// 2.当现金<=50000,每次交1000
	// 计算该人可以经过多少次路口 使用for break方式完成
	var yuan float64 = 1000000
	var count int = 0
	// label01:
	for {
		if yuan > 500000 {
			yuan -= 5.0/100.0*yuan
			count ++
		}
		if yuan <= 500000 && yuan > 1000 {
			yuan -=1000.0
			count ++
		} else if yuan < 1000 {
			break 
		}
	}
	fmt.Printf("可以经过%v次路口",count)
}