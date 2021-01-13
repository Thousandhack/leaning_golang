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

// 查询
/*
	CREATE DATABASE go_demo;
	use go_demo;

	CREATE TABLE `user` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) DEFAULT '',
    `age` INT(11) DEFAULT '0',
    PRIMARY KEY(`id`)
	)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

	插入数据：
		insert into user(name,age) values("test_user",90);
		insert into user(name,age) values("test_two",80);

*/

type user struct {
	id   int
	age  int
	name string
}

// 查询单条数据示例
func queryRowDemo(id int) {
	var u user
	// 写查询单条记录的sql语句
	sqlStr := `select id, name, age from user where id=?`
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	// 下面的 1 为变量，可以进行自己更改
	// 如果不调用Scan，数据库的连接数无法释放
	err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	// 查询第二条数据
	queryRowDemo(2)
}
