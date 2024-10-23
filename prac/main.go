package main

import (
	"fmt"
	"log"
	"strings"
)

// func main() {
// 	var value string

// 	log.Println("Enter the calculate value:")
// 	fmt.Scanln(&value)
// 	operators := "*/+-"
// 	var numberbuilder strings.Builder
// 	// Calculation := operations( 1 ,2 , '+' )
// 	var numbers []int
// 	var valueOperators []rune
// 	for i := 0; i < len(value); {
// 		r, size := utf8.DecodeRuneInString(value[i:])
// 		i += size
// 		fmt.Println(r, "this is codeee")
// 		if strings.ContainsRune(operators, r) {

// 			if numberbuilder.Len() > 0 {
// 				num, _ := strconv.Atoi(numberbuilder.String())
// 				numbers = append(numbers, num)
// 				numberbuilder.Reset()
// 			}
// 			valueOperators = append(valueOperators, r)
// 		} else {
// 			numberbuilder.WriteRune(r)
// 		}
// 	}
// 	if numberbuilder.Len() > 0 {
// 		num, _ := strconv.Atoi(numberbuilder.String())
// 		numbers = append(numbers, num)
// 	}

// 	result := numbers[0]
// 	for i := 0; i < len(valueOperators); i++ {
// 		result = operations(result, numbers[i+1], valueOperators[i])
// 	}

// 	// Output the result
// 	log.Println("Result:", result)
// }
// func operations(a int, b int, operator rune) int {
// 	switch operator {
// 	case '/':
// 		return a / b
// 	case '*':
// 		return a * b
// 	case '+':
// 		return a + b
// 	case '-':
// 		return a - b
// 	default:
// 		return 0
// 	}
// }

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

// func main() {
// 	var value string
// 	log.Println("Enter a value:")
// 	fmt.Scanln(&value)

// 	var numbers []int
// 	var operators []rune
// 	numStr := ""
// 	for _, r := range value {
// 		if isOperator(r) {
// 			num, _ := strconv.Atoi(numStr)
// 			numbers = append(numbers, num)
// 			operators = append(operators, r)
// 			numStr = ""
// 		} else {
// 			numStr += string(r)
// 			log.Println("number string", numStr)
// 		}
// 	}

// 	num, _ := strconv.Atoi(numStr)
// 	numbers = append(numbers, num)

// 	result := numbers[0]
// 	for i, op := range operators {
// 		result = performOperation(result, numbers[i+1], op)
// 	}

// 	log.Println("Result:", result)
// }

// func isOperator(r rune) bool {
// 	return r == '+' || r == '-' || r == '*' || r == '/'
// }

// func performOperation(a, b int, op rune) int {
// 	switch op {
// 	case '+':
// 		return a + b
// 	case '-':
// 		return a - b
// 	case '*':
// 		return a * b
// 	case '/':
// 		return a / b
// 	default:
// 		return 0
// 	}
// }

func main() {
	var expression string
	fmt.Print("Enter the expression: ")
	fmt.Scanln(&expression)

	result := calculate(expression)
	fmt.Printf("Result: %.2f\n", result)
}

func calculate(expr string) float64 {
	expr = strings.ReplaceAll(expr, " ", "")

	var stack []float64
	var num float64
	var sign byte = '+'

	for i := 0; i < len(expr); i++ {
		c := expr[i]
		log.Print(c)
		isDigit := c >= '0' && c <= '9'
		// log.Print(isDigit, "digittt")
		if isDigit {
			num = num*10 + float64(c-'0')
			log.Print(num)

		}
		if !isDigit || i == len(expr)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			sign = c
			num = 0
		}
	}

	var result float64
	for _, v := range stack {
		result += v
	}
	return result
}
