package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func fileWrite(){
	// 没有就创建与追加内容模式
	fileObj,err := os.OpenFile("./log.txt",os.O_CREATE|os.O_APPEND,0666)
	if err != nil{
		fmt.Printf("open file failed,err:%v",err)
		return
	}
	defer fileObj.Close()
	str := "hello 南山\n"
	fileObj.Write([]byte(str))       //写入字节切片数据
	fileObj.WriteString("hello 大沙河") //直接写入字符串数据


}

func fileTRUNC(){
	// 清空文件
	fileObj,err := os.OpenFile("./log.txt",os.O_TRUNC|os.O_WRONLY,0666)
	if err != nil{
		fmt.Printf("open file failed,err:%v",err)
		return
	}
	defer fileObj.Close()
	fileObj.WriteString("清空后再写入的数据") //直接写入字符串数据
}

func fileWritebufio() {
	file, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	//for i := 0; i < 10; i++ {
	//	writer.WriteString("hello宝安\n") //将数据先写入缓存
	//}
	writer.WriteString("hello宝安\n") //将数据先写入缓存
	writer.Flush() //将缓存中的内容写入文件
}

// ioutil.WriteFile
func fileWriteioutil() {
	str := "hello 海上世界"
	err := ioutil.WriteFile("./log.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func main()  {
	// 文件写入操作
	//  os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能。
	/*
	os.O_WRONLY	只写
	os.O_CREATE	创建文件
	os.O_RDONLY	只读
	os.O_RDWR	读写
	os.O_TRUNC	清空
	os.O_APPEND	追加
	*/
	// 打开文件写内容
	//fileWrite()
	// 清空文件
	//fileTRUNC()

	fileWritebufio()
	// ioutil方式写入数据,清空文件内容写入
	fileWriteioutil()

}
