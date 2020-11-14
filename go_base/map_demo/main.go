package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)


func assignSort() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

func main()  {
	var m1 map[string]int   // 还没有初始化（没有再内存中开闭内存空间）
	m1 = make(map[string]int,10) // 要估算好该map容量，避免在程序运行期间再动态扩容（这样执行效率更高）
	m1["时光"] = 18
	m1["阳光"] = 25

	fmt.Println(m1)
	fmt.Println(m1["时光"])
	value,ok := m1["hsz"]
	// ok 是一个bool类型，ok是约定成俗的
	if !ok {
		fmt.Println("查不到此key")
	} else {
		fmt.Println(value)
	}
	// map 遍历
	for k,v := range m1{
		fmt.Println(k,v)
	}
	// 只打印键
	for k := range m1{
		fmt.Println(k)
	}
	// 只打印值
	for _,v := range m1{
		fmt.Println(v)
	}
	// 删除
	delete(m1,"时光")
	fmt.Println(m1)

	// 指定排序输出
	assignSort()

}
