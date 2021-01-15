package main

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql" // init()

var db *sql.DB // 是一个连接池对象

type user struct {
	id   int
	age  int
	name string
}

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

// 预处理查询
func prepareQueryDemo() {
	sqlStr := `select id,name,age from user where id > ?`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return
	}
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scanf failed,err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 预处理插入示例
func prepareInsertDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	stmt, err := db.Prepare(sqlStr)  // 把SQL语句先发给MySQL预处理一下
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	// 插入一条
	//_, err = stmt.Exec("test_小王子", 18)
	//if err != nil {
	//	fmt.Printf("insert failed, err:%v\n", err)
	//	return
	//}

	// 插入多条数据
	// 注意map是无序的
	var m = map[string]int{
		"test_xx": 28,
		"test_ff": 23,
		"test_vv": 233,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}
	fmt.Println("insert success.")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	// 预处理查询
	prepareQueryDemo()
	// 预处理插入  一条数据与多条数据
	prepareInsertDemo()

}
