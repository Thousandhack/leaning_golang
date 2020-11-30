package main

import (
	"fmt"
	"os"
)

func demo(){
	// 以下功能是替换内容
	fileObj,err := os.OpenFile("./text.txt",os.O_RDWR,0644)
	if err != nil{
		fmt.Printf("open file failed,err:%v\n",err)
		return
	}
	defer fileObj.Close()
	fileObj.Seek(3,0)  // 光标移动
	var s []byte
	s = []byte{'c','a','d'}
	fileObj.Write(s)
	var ret [3]byte // 因为文件中的是中文，所以注意，一个中文3个字节 ，如果字节数不够一个汉字会打印出乱码
	n,err := fileObj.Read(ret[:])
	if err != nil{
		fmt.Printf(" read from file failed,err:%v\n",err)
		return
	}
	fmt.Println(string(ret[:n]))

}

func insertContent(){
	fileObj,err := os.OpenFile("./text.txt",os.O_RDWR,0644)
	if err != nil{
		fmt.Printf("open file failed,err:%v\n",err)
		return
	}
	defer fileObj.Close()
	// 因为没有办法直接在文件中插入内容，所以要借助一个临时文件
	temFile,err := os.OpenFile("./temp.txt",os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0644)
	if err != nil{
		fmt.Printf("open file failed,err:%v\n",err)
		return
	}
	defer temFile.Close()
	// 读原文件写入原始文件
	// 读取文件写入临时文件
	var ret [3]byte
	n,err := fileObj.Read(ret[:])
	if err != nil{
		fmt.Printf("open file failed,err:%v\n",err)
		return
	}
	// 写入临时文件
	temFile.Write(ret[:n])
	// 再写入要插入的内容
	var s []byte
	s = []byte{'c','a','b'}
	temFile.Write(s)


}

func main()  {
	//demo()
	insertContent()
}
