package main

import "fmt"

type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
// dog 为接受者
// 约定成俗变量的首字母 类似python中的self
func (d dog) wang() {
	fmt.Printf("%s汪汪汪~", d.name)
}

func main() {
	d := newDog("逗逗")
	d.wang()
}
