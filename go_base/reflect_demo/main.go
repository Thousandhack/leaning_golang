package main

import (
	"encoding/json"
	"fmt"
)

// json
type person struct {
	Name string `json:"name"`
	Age  int `json:"age"`
}

func main() {
	str := `{"name":"hsz","age":9000}`
	var p person
	// 这里面就是用了反射的原理
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name, p.Age)
}
