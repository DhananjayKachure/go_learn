package main

import "fmt"

func main() {

	value := []int{1, 2, 3}

	printvalue(value)

}

func printvalue[T string | int](params []T) {
	fmt.Println(params)
}
