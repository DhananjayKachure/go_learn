package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	//time format
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	//time.Date(year , time.Month , date , hours , minute ,seconds , seconds , standardtime)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))

}
