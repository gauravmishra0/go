package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("TIME")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 onday"))

	createdDate := time.Date(2020, time.April, 7, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
