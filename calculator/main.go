package main

import (
	"fmt"
)

func main() {
	var expression string
	fmt.Print("Enter the expression: ")
	fmt.Scanln(&expression)

	result := calculate(expression)
	fmt.Printf("Result: %v\n", result)
}

func calculate(expr string) float64 {
	var operations []float64
	var num float64
	var sign byte = '+'

	for i := 0; i < len(expr); i++ {
		c := expr[i]
		isDigit := c >= '0' && c <= '9'
		isAlpha := (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
		if isAlpha {
			fmt.Println("Error: Expression contains alphabetic characters.")
			return 0
		}
		if isDigit {
			num = num*10 + float64(c-'0')
		}
		if (!isDigit && c != ' ') || i == len(expr)-1 {
			switch sign {
			case '+':
				operations = append(operations, num)
			case '-':
				operations = append(operations, -num)
			case '*':
				operations[len(operations)-1] *= num
			case '/':
				operations[len(operations)-1] /= num
			}
			sign = c
			num = 0
		}
	}
	var result float64
	for _, v := range operations {
		result += v
	}

	return result
}
