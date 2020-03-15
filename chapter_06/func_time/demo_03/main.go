package main

import "fmt"
import "time"

func main() {
	now := time.Now()
	// Unix 和 UnixNamo
	fmt.Printf("uinx时间戳=%v \nunixnano时间戳=%v",now.Unix(),now.UnixNano())

}