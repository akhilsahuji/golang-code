package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to study of time golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2000, time.March, 06, 23, 52, 0, 0, time.UTC)
	fmt.Println(createDate)

	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))

}
