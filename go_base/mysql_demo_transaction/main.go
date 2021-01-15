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

func transactionDemo() {
    // 1. 开启事务
    tx, err := db.Begin()
    if err != nil {
        fmt.Printf("begin failed,err:%v\n", err)
        return
    }
    // 执行多个sql操作
    sqlStr1 := `update user set age=age-2 where id=1`
	sqlStr2 := `update bbb user set age=age+2 where id=100`  //
    //sqlStr2 := `update bbb user set age=age+2 where id=100`  // 用此句sql会回滚，两个数据都不会更新
    // 执行sql1
    _, err = tx.Exec(sqlStr1)
    if err != nil {
        // 要回滚
        tx.Rollback()
		fmt.Printf("第一跳数据修改出错,err:%v\n",err)
        return
    }
    // 执行sql2
    _, err = tx.Exec(sqlStr2)
    if err != nil {
        // 要回滚
        tx.Rollback()
		fmt.Printf("第二跳数据修改出错,err:%v\n",err)
        return
    }
    // 执行提交
    err = tx.Commit() // 提交事务
    if err != nil {
        tx.Rollback()
        fmt.Printf("提交出错回滚!,err:%v\n",err)
    }
}

func main() {
    err := initDB()
    if err != nil {
        fmt.Printf("init DB failed,err:%v\n", err)
        return
    }

    transactionDemo()

}
