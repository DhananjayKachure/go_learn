package main

import (
	"fmt"
	"strconv"
)

func main() {

	a := "5"
	b := 4
	number, _ := strconv.Atoi(a)
	result := add(number, b)
	fmt.Println("this is number", result)
}

func add(value1 int, value2 int) int {
	return value1 + value2
}

func subtract(value1 int, value2 int) int {
	return value1 - value2
}

func multiply(value1 int, value2 int) int {
	return value1 * value2
}

func divide(value1 int, value2 int) int {
	return value1 / value2
}
