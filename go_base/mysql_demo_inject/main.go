package main

import (
    "fmt"
)
import "github.com/jmoiron/sqlx"
import _ "github.com/go-sql-driver/mysql" // init()

var db *sqlx.DB

type user struct {
    ID   int
    Age  int
    Name string
}

func initDB() (err error) {
    database := "root:123456@tcp(127.0.0.1:3306)/go_demo?charset=utf8mb4&parseTime=True"
    // 也可以使用MustConnect连接不成功就panic
    db, err = sqlx.Connect("mysql", database)
    if err != nil {
        fmt.Printf("connect DB failed, err:%v\n", err)
        return
    }
    //db.SetMaxOpenConns(20)
    //db.SetMaxIdleConns(10)
    return
}

// SQL注入
func sqlInjectDemo(name string) {
    sqlStr := fmt.Sprintf("select id,name,age from user where name='%s'", name)
    fmt.Printf("SQL:%s\n", sqlStr)

    var users []user
    err := db.Select(&users, sqlStr)
    if err != nil {
        fmt.Printf("exec failed,err:%v\n", err)
        return
    }
    for _, u := range users {
        fmt.Printf("user:%#v\n", u)
    }
    // SQL注入的几种示例
}

func main() {
    err := initDB()
    if err != nil {
        fmt.Printf("init DB failed,err:%v\n", err)
        return
    }
    // 正常插入
    sqlInjectDemo("test_user")
    fmt.Println("================================================")
    // sql注入的问题的例子
    sqlInjectDemo("test_user' or 1=1 #")
    fmt.Println("================================================")
    sqlInjectDemo("xxx' union select * from user #")
    fmt.Println("================================================")
    sqlInjectDemo("xxx' and (select count(*) from user) <10 #")

}

// 解决方法一般是用预编译