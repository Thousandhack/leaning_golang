package main

import "fmt"
import "sort"

func main(){

	// map的排序
	map1 := make(map[int] int)
	map1[10] = 100
	map1[1] = 10
	map1[5] = 77
	map1[7] = 21
	map1[55] = 21
	fmt.Println(map1)

	// for k,v := range map1 {
	// 	fmt.Printf("k=%v,v=%v\n",k,v)
	// }



	// 如果按照map的key的顺序排序输出
	// 1. 先将map的key 放入到 切片中
	// 2. 对切片排序
	// 3. 遍历切片 然后按照key值来输出map的值
	var keys []int
	for k,_ := range map1 {
		keys = append(keys,k)
	}

	sort.Ints(keys)
	fmt.Println(keys)  // map 的key值排序好了

	for _,key := range keys{
		fmt.Printf("key=%v,value=%v\n",key,map1[key])
	}
}
