package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch for go")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("move 1")
	case 2:
		fmt.Println("move 2")
	case 3:
		fmt.Println("move 3")
		fallthrough
	case 4:
		fmt.Println("move 4")
	case 5:
		fmt.Println("move 5")
	case 6:
		fmt.Println("move 6")
	default:
		fmt.Println("move 0")

	}

}
