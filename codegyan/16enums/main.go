package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func main() {
	confirmbookingStatus(Confirmed)
}

func confirmbookingStatus(status OrderStatus) {
	fmt.Println("ur statusss", status)
}
