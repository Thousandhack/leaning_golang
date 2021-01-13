package main

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql" // init()

var db *sql.DB // 是一个连接池对象

func initDB() (err error) {
	// 数据库信息
	database := "root:123456@tcp(127.0.0.1:3306)/go_demo"
	// 连接数据库
	// 这边注意，现在改为了全局的db
	db, err = sql.Open("mysql", database) // 不会效验用户名和密码是否正确，校验的是格式
	if err != nil {
		fmt.Printf("database %s invalid,err:%v\n", database, err)
		return
	}
	err = db.Ping() // 校验连接数据库是否成功
	if err != nil {
		fmt.Printf("open %s failed,err:%v\n", database, err)
		return
	}
	fmt.Println("连接数据库成功!")
	return
}

// 插入单条数据示例
func insertDemo(name string, age int) {
	// 1.写SQL语句
	sqlStr := `insert into user(name,age) values(?,?)`
	// 2.exec
	ret, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed,err:%v", err)
	}
	// 如果插入数据的操作，能够拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Println(id)
	fmt.Println(ret)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	// 插入数据函数
	insertDemo("test_night", 29)
}
