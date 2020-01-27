// package main

// import "fmt"

// func main() {
// 	// (1)写一个程序，获取一个int变量num的地址,并显示到终端
// 	// (2)将num的地址赋给指针ptr,并通过ptr去修改num的值
// 	var num int = 9
// 	fmt.Printf("num address=%v\n", &num)

// 	var ptr *int 
// 	ptr = &num
// 	*ptr = 10 //这里修改时，会到num的值变化，指针地址对应相应的值
// 	fmt.Println("num =" , num)
// }

package main
import "fmt"

func main(){
	//
	var a int = 300
	var b int = 400
	var ptr *int = &a // ptr指针变量为a的地址
	*ptr = 100 // a = 100
	ptr = &b  // ptr赋值为b的地址
	*ptr = 200
	fmt.Printf("a=%d,b=%d,*ptr=%d",a,b,*ptr)
	// a = 100 b = 100  *ptr = 200
}