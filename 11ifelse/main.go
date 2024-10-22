package main

import "fmt"

func main() {
	fmt.Println("if elseeee")
	logicCount := 10
	var result string
	if logicCount < 10 {
		result = "its less then 10"
	} else if logicCount > 10 {
		result = "its greater then 10"
	} else {
		result = "equals to 10"
	}
	fmt.Println(result)
	if num := 3; num < 10 {
		fmt.Println("its less then 10")
	} else {
		fmt.Println("its greater 10")
	}
}
