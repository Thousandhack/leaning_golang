package main

import (
	"fmt"
	"os"
	"time"
)

type logger interface {
	consoleLog()
	fileLog()
}

type User struct {
	username string
	password string
}

func (u User) consoleLog() {
	t := time.Now()
	data := "用户创建成功！ 用户名为：" + fmt.Sprintf("%s\n", u.username) + "密码是：" + u.password + "\n" + fmt.Sprintf("创建完成时间：%d-%d-%d %d:%d:%d\n", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Println(data)
}

func (u User) fileLog() {
	t := time.Now()
	// |os.O_RDWR
	file, err := os.OpenFile("./10_consoleLog_fileLog/"+u.username+".txt", os.O_APPEND|os.O_CREATE, 0766) // 如果有这个文件就打开没有就新建
	if err != nil {
		fmt.Println(err)
	}
	data := "用户创建成功！ 用户名为：" + fmt.Sprintf("%s\n", u.username) + "密码是：" + u.password + "\n" + fmt.Sprintf("创建完成时间：%d-%d-%d %d:%d:%d\n", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	_, _ = file.WriteString(data)
	_ = file.Close()
}

func newUser(username, password string) User {
	return User{
		username: username,
		password: password,
	}
}

func main() {
	//
	var (
		username string
		password string
	)
	fmt.Print("请输入用户名：")
	_, err := fmt.Scan(&username)
	fmt.Print("请输入一个密码：")
	_, err = fmt.Scan(&password)
	if err != nil {
		fmt.Println("输入错误！！ERROR:", err)
	}
	u := newUser(username, password)
	u.consoleLog()
	u.fileLog()

}

/*
	O_RDONLY：只读模式(read-only)
	O_WRONLY：只写模式(write-only)
	O_RDWR：读写模式(read-write)
	O_APPEND：追加模式(append)
	O_CREATE：文件不存在就创建(create a new file if none exists.)
	O_EXCL：与 O_CREATE 一起用，构成一个新建文件的功能，它要求文件必须不存在(used with O_CREATE, file must not exist)
	O_SYNC：同步方式打开，即不使用缓存，直接写入硬盘
	O_TRUNC：打开并清空文件
*/