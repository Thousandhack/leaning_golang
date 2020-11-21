package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// 结构体与json

// 1.序列化:把Go语言中的结构体变量 --> json格式的字符串
// 2.反序列化： json格式的字符串  --> Go语言中能够识别的结构体变量
type person struct {
	Name string `json:"name" db:"name"`   // json情况下使用字段
	Age int     `json:"age"`
}


func main(){
	p1 := person{
		Name: "zero",
		Age:  18,
	}
	b,err := json.Marshal(p1)  // 转成字节流
	if err !=nil{
		fmt.Println("")
		return
	}
	fmt.Printf("%#v\n",string(b)) // 字节流可以强转为字符串
	fmt.Printf("%v\n",string(b))
	// 反序列化
	str := `{"name":"one","age":19}`
	var p2 person
	er := json.Unmarshal([]byte(str),&p2)
	if er !=nil{
		errors.New("unmarshal error")
	}
	fmt.Printf("%v\n",p2)
	fmt.Println(p2.Name,p2.Age)
}
