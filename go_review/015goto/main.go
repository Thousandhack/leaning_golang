package main

import "fmt"

/*
goto语句通过标签进行代码间的无条件跳转。
goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。
*/
func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签 breakTag为指定的标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	// 标签
breakTag:
	fmt.Println("结束for循环")
	return

}

func main() {
	gotoDemo2()
}
