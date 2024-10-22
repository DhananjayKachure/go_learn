package main

import "fmt"

func main() {
	result := addone(1, 2)
	fmt.Println(result)
	prores, message := proFunction(1, 2, 3, 4, 5)
	fmt.Println("result:", prores, message)
}

func addone(value1 int, value2 int) int {
	return value1 + value2
}

func proFunction(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "heyyyyyyyy"
}
