package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 97
	result := strconv.Itoa(number)

	fmt.Printf("%T\n", number)
	fmt.Println(result)
	fmt.Printf("%T\n", result)
}
