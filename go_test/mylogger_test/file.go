package mylogger_test

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件中写入日志，并在文件日志一定大小后，进行切割

type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件保存的文件名
	fileObj     *os.File
	errfileObj  *os.File
	errFileName string // 记录error以下的级别的日志
	maxFileSize int64  // 最大的文件大小
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() // 按照文件路径和文件名将文件打开

	return fl
}

func (f *FileLogger) initFile() (error) {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v", err)
		return err
	}
	// 日志文件都已经打开了
	f.fileObj = fileObj
	f.errfileObj = errFileObj
	return nil
}

func (f *FileLogger) enable(loglevel LogLevel) bool {
	return loglevel >= f.Level
}

func (f *FileLogger) checkFileSize(file *os.File) bool {
	// *os.File 是一个结构体指针
	// 获取文件对象的详细信息
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return false
	}
	//fmt.Printf("文件大小是：%d\n Byte",fileInfo.Size())
	// 如果当前文件大小大于等于日志文件的最大值，就应该返回True
	return fileInfo.Size() >= f.maxFileSize

}

func (f *FileLogger) splitFileLog(file *os.File) (*os.File, error) {
	// 日志需要切割
	// 1.关闭当前的日志文件
	f.fileObj.Close()
	// 2.rename 进行备份
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return nil, err
	}
	oldLogFileName := path.Join(f.filePath, fileInfo.Name()) // 拿到当前日志的完整路径
	newLogFileName := fmt.Sprintf("%s_%s.bak_%s", f.filePath, f.fileName, nowStr)
	os.Rename(oldLogFileName, newLogFileName)
	// 3.打开一个新的文件
	fileObj, err := os.OpenFile(oldLogFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return nil, err
	}
	// 4.将打开的新日志文件对象赋值给f.fileObj
	return fileObj, nil
}

// 公用的写入文件的方法
func (f FileLogger) fileLog(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3) // 拿到谁调用的这一行的功能
		if f.checkFileSize(f.fileObj) {
			newFile,err := f.splitFileLog(f.fileObj)  // 日志文件
			if er
		}
		now.Format("2006-01-02 15:04:05")
		fmt.Fprintf(f.fileObj, "%s [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogStringLevel(lv), funcName, fileName, lineNo, msg)
		if lv > ERROR {
			// 如果要记录的日志大于等于ERROR级别，还要在err日志文件中再记录一遍
			fmt.Fprintf(f.errfileObj, "%s [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogStringLevel(lv), funcName, fileName, lineNo, msg)
		}
	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	if f.enable(DEBUG) {
		f.fileLog(DEBUG, format, a...)
	}
}

func (f FileLogger) Trance(format string, a ...interface{}) {
	if f.enable(TRANCE) {
		f.fileLog(TRANCE, format, a...)
	}
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	if f.enable(INFO) {
		f.fileLog(INFO, format, a...)
	}
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
	if f.enable(WARNING) {
		f.fileLog(WARNING, format, a...)
	}
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	if f.enable(ERROR) {
		f.fileLog(ERROR, format, a...)
	}
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	if f.enable(FATAL) {
		f.fileLog(FATAL, format, a...)
	}
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errfileObj.Close()
}
