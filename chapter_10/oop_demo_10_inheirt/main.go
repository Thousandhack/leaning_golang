package main

import "fmt"

type A struct{
   Name string
   age int
}

type B struct{
   Name string
   Score float64
}

type C struct {
   A
   B
}

type D struct{
   a A // 有名结构体
}

func main(){
   var c  C
   // c.Name = "zero"  // error
   c.A.Name = "zero"
   fmt.Println(c.A.Name)  // 需要写明白继承哪个结构体的字段的

   var d D
   d.a.Name = "one"
   fmt.Println(d)
   fmt.Println(d.a.Name)  // 有名结构体的使用
}
