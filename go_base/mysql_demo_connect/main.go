package main

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql" // init()

/*
	主要实现mysql连接数据库的成功
*/
func main() {
	// 信息数据库
	database := "root:123456@tcp(127.0.0.1:3306)/go_demo"
	// 连接数据库
	db, err := sql.Open("mysql", database) // 不会效验用户名和密码是否正确，校验的是格式
	if err != nil {
		fmt.Printf("database %s invalid,err:%v\n", database, err)
		return
	}
	err = db.Ping() // 校验连接数据库是否成功
	if err != nil {
		fmt.Printf("open %s failed,err:%v\n", database, err)
		return
	}
	fmt.Println("连接数据库成功")
}
