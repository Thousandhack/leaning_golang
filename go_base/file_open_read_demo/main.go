package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func fileOpen() {
	// 只读方式打开当前目录下的log.txt文件
	file, err := os.Open("./log.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	// 关闭文件
	defer file.Close()
	// 读完了关闭文件
	fmt.Println("打开了")
}

func fileRead() {
	// 只读方式打开当前目录下的log.txt文件
	file, err := os.Open("./log.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
	// 使用Read方法读取数据
	var tmp = make([]byte, 128) // 只读了128kb的数据，有可能会没有读完
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))
}

func fileAllRead() {
	// 只读方式打开当前目录下的log.txt文件
	file, err := os.Open("./log.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
	// 循环读取文件
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

// bufio按行读取示例
func fileReadTypebufio() {
	file, err := os.Open("./log.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

// ioutil.ReadFile读取整个文件
func fileReadTypeioutil() {
	content, err := ioutil.ReadFile("./log.txt")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	// 打开文件
	fileOpen()
	// 读文件
	fileRead()
	// 循环读取所有文件
	fmt.Println("循环读取所有文件")
	fileAllRead()
	// bufio按行读取示例  读一行一行的
	fmt.Println("bufio按行读取示例")
	fileReadTypebufio()

	// ioutil.ReadFile读取整个文件
	fmt.Println("ioutil.ReadFile读取整个文件内容")
	fileReadTypeioutil()
}
