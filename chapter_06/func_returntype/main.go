package main

import "fmt"

func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2

	sub = n1 - n2

	return
}

func main() {
	//
	a, b := getSumAndSub(3, 2)
	fmt.Println(a, b)

}
