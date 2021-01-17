package main

import "fmt"
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

func queryDemo() {
    var u user
    sqlStr1 := `select id,name,age from user where id=1`
    db.Get(&u, sqlStr1)
    fmt.Printf("u:%#v\n", u)
}

func queryListDemo() {
    var userList []user
    sqlStr2 := `select id,name,age from user`
    err := db.Select(&userList, sqlStr2)
    if err != nil {
        fmt.Printf("select failed,err:%v\n", err)
        return
    }
    fmt.Printf("userList:%#v\n", userList)
}


// 插入数据
func insertRowDemo() {
    sqlStr := "insert into user(name, age) values (?,?)"
    ret, err := db.Exec(sqlStr, "test_iix", 19)
    if err != nil {
        fmt.Printf("insert failed, err:%v\n", err)
        return
    }
    theID, err := ret.LastInsertId() // 新插入数据的id
    if err != nil {
        fmt.Printf("get lastinsert ID failed, err:%v\n", err)
        return
    }
    fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo() {
    sqlStr := "update user set age=? where id = ?"
    ret, err := db.Exec(sqlStr, 39, 18)
    if err != nil {
        fmt.Printf("update failed, err:%v\n", err)
        return
    }
    n, err := ret.RowsAffected() // 操作影响的行数
    if err != nil {
        fmt.Printf("get RowsAffected failed, err:%v\n", err)
        return
    }
    fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
    sqlStr := "delete from user where id = ?"
    ret, err := db.Exec(sqlStr, 11)
    if err != nil {
        fmt.Printf("delete failed, err:%v\n", err)
        return
    }
    n, err := ret.RowsAffected() // 操作影响的行数
    if err != nil {
        fmt.Printf("get RowsAffected failed, err:%v\n", err)
        return
    }
    fmt.Printf("delete success, affected rows:%d\n", n)
}

func main() {

    err := initDB()
    if err != nil {
        fmt.Printf("init DB failed,err:%v\n", err)
        return
    }
    // 查询单条数据
    queryDemo()
    // 查询多条数据
    queryListDemo()
    // 插入一条数据
    //insertRowDemo()
    // 更新数据
    updateRowDemo()
    // 删除数据
    deleteRowDemo()

}
