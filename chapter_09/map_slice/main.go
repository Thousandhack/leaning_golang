package main

import "fmt"


func main(){
	// 1.声明一个map切片
	var monsters []map[string]string
	monsters = make ([]map[string]string,2)  // 只给了两个空间
	// 2.添加数据
	if monsters[0] == nil {
		monsters[0] = make(map[string]string,2)
		monsters[0]["name"] = "悟空"
		monsters[0] ["age"] = "550"
	}
	
	if monsters[1] == nil {
		monsters[1] = make(map[string]string,2)
		monsters[1]["name"] = "八戒"
		monsters[1] ["age"] = "490"
	}


	fmt.Println(monsters)

	// 如果想要数据增加使用append 函数动态增加
	// 先定义monsters信息
	newMonster := map[string]string{
		"name" : "沙和尚",
		"age":"466",
	}

	monsters = append(monsters,newMonster)
	// map的里面的数据就新增成功了
	fmt.Println(monsters)


}