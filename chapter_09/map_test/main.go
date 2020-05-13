package main

import "fmt"

func modifyUser(users map[string]map[string]string,name string){
	if users[name] != nil{
		users[name]["password"] = "888888"
	} else {
		users[name] = make(map[string]string,2)
		users[name]["password"] = "666666"
		users[name]["nick_name"] = "hahaha"
	}
}

func main()  {
	// 使用 map[string]map[string]string 的map类型
	// key 表示用户名，是唯一的
	// 如果存在用户名hsz ，将密码修改为888888 
	// 否则将用户信息添加
	users := make(map[string]map[string]string,10)
	users["zero"] = make(map[string]string,2)
	users["zero"]["password"] = "213131"
	users["zero"]["nick_name"] = "000z"
	modifyUser(users,"tome")
	modifyUser(users,"zero")

	fmt.Println(users)
}