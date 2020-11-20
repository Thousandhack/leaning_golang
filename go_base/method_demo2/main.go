package method_demo2

import "fmt"

// 给自定义类型加方法

type  myInt int

func (m myInt)hello()  {
	fmt.Println("我是一个自己的int类型")
	fmt.Printf("%T\n",m)
}

func main()  {
	//
	m := myInt(100)
	m.hello()
}
