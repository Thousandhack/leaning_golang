package main


import "fmt"
import "encoding/json"

type Monster struct{
	Name string `json:"name"`   // `json:"name"` 这个就是结构体的标签tag
	Age int	`json:"age"`        // 效果是可以首字母大写的变量可以在json返回小写
	Skill string `json:"skill"`
}

func main(){

	// 1.创建一个Monster变量
	monster := Monster{"ZEOR",25,"coding go"}

	// 2.将monster变量序列化为json格式的字符串
	//json.Marshal 函数中使用到了反射
	jsononster,err := json.Marshal(monster)
	if err != nil{
		fmt.Println("json处理错误",err)
	}
	fmt.Println("jsononster:",string(jsononster))  // jsononster数据类型为bytes需求强转为string
}