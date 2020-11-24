package main

import (
	"fmt"
	"gotest01/go_base/package_demo/calc_demo"  // 到目录下就可以了，然后目录下每个导入package都是最后一层的目录名字
)

func main()  {
	res := calc_demo.Add(5,7)
	fmt.Println(res)
}
