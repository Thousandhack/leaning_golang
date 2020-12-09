package main

import (
	"gotest01/go_test/mylogger_test"
	"time"
)

//  可以往不同的输出位置记录日志
//  日志分为五种级别
/*
	1.debug
	2.Trance
	3.Info
	4.Warning
	5.Error
	6.Fatal
*/
// 日志要支持开关控制  开发环境什么都输出，但是上线之后只有输出INFO
// 日志的完整日志记录要有时间，行号，文件名，日志级别，日志信息
// 日志文件要切割
	// 1.按文件大小切割
		// 1.每次记录日志之前都判断一下当前写的这个文件的文件的大小
	// 2.按时期切割
		// 在日志结构体中设置一个字段记录上一次切割的小时数
		// 在写日志之前检查一下当前的小时数和之前保存的是否一致，不一致就要切割
//
// 第三方日志库：
//如logrus、zap等

var log mylogger_test.Logger  // 声明一个全局的接口变量

func main() {
	//fileObj,err := os.OpenFile("./log.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	//if err != nil{
	//	fmt.Printf("open file failed,err:%v\n",err)
	//	return
	//}
	//log.SetOutput(fileObj)
	//for {
	//	log.Println("这是一条测试的日志")
	//	time.Sleep(time.Second*3)
	//}

	log = mylogger_test.NewLog("debug")  // 设置将日志写入文件的级别

	//log = mylogger_test.NewFileLogger("Info", "./", "hsz.log", 10*1024)
	// 实现可以穿字符串，也可以传格式化变量值
	for {
		name := "hsz"
		log.Debug("这是一条debug级别的日志name:%s", name)
		log.Info("这是一条info级别的日志name:%s", name)
		log.Trace("这是一条Trance级别的日志")
		log.Warning("这是一条Warning级别的日志")

		log.Error("这是一条error级别的日志")
		log.Fatal("这是一条Fatal级别的日志")
		time.Sleep(time.Second * 1)
	}

}
