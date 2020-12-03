package main

import (
	"fmt"
	"log"
	"os"
	"time"
)


// 往日志文件写日志的简单例子
func main() {
	fileObj,err := os.OpenFile("./log.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil{
		fmt.Printf("open file failed,err:%v\n",err)
		return
	}
	log.SetOutput(fileObj)
	for {
		log.Println("这是一条测试的日志")
		time.Sleep(time.Second*3)
	}
	//for {
	//	name := "hsz"
	//	log.Debug("这是一条debug级别的日志name:%s",name)
	//	log.Info("这是一条info级别的日志name:%s",name)
	//	log.Trance("这是一条Trance级别的日志")
	//	log.Warning("这是一条Warning级别的日志")
	//
	//	log.Error("这是一条error级别的日志")
	//	log.Fatal("这是一条Fatal级别的日志")
	//	time.Sleep(time.Second * 3)
	//}
}
