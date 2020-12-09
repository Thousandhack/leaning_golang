package mylogger_test

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

// Logger 接口
type Logger interface {
	Debug(format string,a...interface{})
	Info(format string,a...interface{})
	Trace(format string,a...interface{})
	Warning(format string,a...interface{})
	Error(format string,a...interface{})
	Fatal(format string,a...interface{})

}

// 往终端写日志相关内容
type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRANCE
	INFO
	WARNING
	ERROR
	FATAL
)


func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRANCE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err

	}

}

func getLogStringLevel(lv LogLevel) string {
	//s = strings.ToLower(s)
	switch lv {
	case  DEBUG:
		return "debug"
	case TRANCE:
		return "trace"
	case INFO:
		return "info"
	case WARNING:
		return "warning"
	case ERROR:
		return "error"
	case FATAL:
		return "fatal"
	default:
		//err := errors.New("无效的日志级别")
		return "unknown"

	}

}

func getInfo(skip int)(funcName,fileName string,lineNo int){
	pc,file,lineNo,ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() fiald\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
	//fmt.Println(funcName)
	//fmt.Println(file)
	//fmt.Println(path.Base(file))  // 获取当前文件名
	//fmt.Println(line)
}

