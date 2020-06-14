package main

import "fmt"


type Calcuator struct{

}


// 实现方法一
func (c *Calcuator) GetAdd(num1 float64,num2 float64) (float64) {
	return num1 + num2
} 

func (c *Calcuator) GetSubtract(num1 float64,num2 float64) (float64) {
	return num1 - num2
} 

func (c *Calcuator) GetMultiply(num1 float64,num2 float64) (float64) {
	return num1 * num2
} 

func (c *Calcuator) GetDivide(num1 float64,num2 float64) (float64) {
	return num1 / num2
} 


// 方法2
func (c *Calcuator) GetResult(num1 float64,num2 float64, key string) (float64){
	fmt.Println(key)
	// return num1 + num2
	if key == "+"{
		return num1 + num2
	} else if key == "-" {
		return num1 - num2
	} else if key == "*" {
		return num1 * num2
	} else if key == "/" {
		return num1 / num2
	} else {
		fmt.Println("error")
		return 0
	}
			
}

func main(){
	// 定义一个计算器结构体（Calcuator）
	// 实现加减乘除四个功能
	// 实现法一：用四个方法完成
	// 实现法二：一个方法完成
	var c Calcuator
	addRes := c.GetAdd(1.00,2.00)
	fmt.Println(addRes)
	subtractRes := c.GetSubtract(1.00,2.00)
	fmt.Println(subtractRes)
	multiplyRes := c.GetMultiply(1.00,2.00)
	fmt.Println(multiplyRes)
	divideRes := c.GetDivide(1.00,2.00)
	fmt.Println(divideRes)

	resValue := c.GetResult(1.00,2.00,"+")
	fmt.Println(resValue)

}

