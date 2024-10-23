package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	var value string

	log.Println("Enter the calculate value:")
	fmt.Scanln(&value)
	operators := "*/+-"
	var numberbuilder strings.Builder
	// Calculation := operations( 1 ,2 , '+' )
	var numbers []int
	var valueOperators []rune
	for i := 0; i < len(value); {
		r, size := utf8.DecodeRuneInString(value[i:])
		i += size
		// log.Println(string(r))

		if strings.ContainsRune(operators, r) {
			if numberbuilder.Len() > 0 {
				num, _ := strconv.Atoi(numberbuilder.String())
				log.Println(num)
				numbers = append(numbers, num)
				numberbuilder.Reset()
			}
			valueOperators = append(valueOperators, r)
		} else {
			fmt.Println(numberbuilder.WriteRune(r))
		}
	}
	if numberbuilder.Len() > 0 {
		num, _ := strconv.Atoi(numberbuilder.String())
		numbers = append(numbers, num)
	}

	// Perform the calculations
	result := numbers[0] // Initialize with the first number
	for i := 0; i < len(valueOperators); i++ {
		result = operations(result, numbers[i+1], valueOperators[i])
	}

	// Output the result
	log.Println("Result:", result)
}
func operations(a int, b int, operator rune) int {
	switch operator {
	case '/':
		return a / b
	case '*':
		return a * b
	case '+':
		return a + b
	case '-':
		return a - b
	default:
		return 0
	}
}

// a := "5"
// b := 4
// number, _ := strconv.Atoi(a)
// result := add(number, b)
// fmt.Println("this is number", result)

// func add(value1 int, value2 int) int {
// 	return value1 + value2
// }

// func subtract(value1 int, value2 int) int {
// 	return value1 - value2
// }

// func multiply(value1 int, value2 int) int {
// 	return value1 * value2
// }

// func divide(value1 int, value2 int) int {
// 	return value1 / value2
// }
