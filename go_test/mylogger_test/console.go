package mylogger_test

import (
	"fmt"
	"time"
)


// 结构体
type ConsoleLogger struct {
	Level LogLevel
}

// 构造函数
func NewLog(levelStr string) ConsoleLogger  {
	level, err := parseLogLevel(levelStr) // 出入日志级别后进行解析成相应的日志级别
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (l ConsoleLogger) enable(loglevel LogLevel) bool {
	return  loglevel >= l.Level
}


// 公用的log函数
func (c ConsoleLogger)log(lv LogLevel,format string,a ...interface{}){
	if c.enable(lv) {
	msg := fmt.Sprintf(format,a...)
	now := time.Now()
	funcName,fileName,lineNo := getInfo(3)  // 拿到谁调用的这一行的功能
	now.Format("2006-01-02 15:04:05")
	fmt.Printf("%s [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"),getLogStringLevel(lv),funcName,fileName,lineNo,msg)
	}
}

func (c ConsoleLogger) Debug(format string,a ...interface{}) {
	c.log(DEBUG,format,a...)
}

func (c ConsoleLogger) Trace(format string,a ...interface{}) {
	c.log(TRANCE,format,a...)
}

func (c ConsoleLogger) Info(format string,a ...interface{}) {
	c.log(INFO,format,a...)
}

func (c ConsoleLogger) Warning(format string,a ...interface{}) {
	c.log(WARNING,format,a...)
}

func (c ConsoleLogger) Error(format string,a ...interface{}) {
	c.log(ERROR,format,a...)
}

func (c ConsoleLogger) Fatal(format string,a ...interface{}) {
	c.log(FATAL,format,a...)
}

