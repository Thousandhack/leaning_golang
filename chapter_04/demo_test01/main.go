package main

import (
	"fmt"
)

func main() {
	// 假如还有97天，问：**个星期零**天
	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Printf("%d个星期零%d天\n", week, day)
}
