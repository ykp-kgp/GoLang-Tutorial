package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time study")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006"))

	createdDate := time.Date(2020, time.April, 9, 21, 21, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
