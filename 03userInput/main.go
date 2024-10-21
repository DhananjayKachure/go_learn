package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("give ratingg to me")
	// comma ok || err err
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating: ", input)
	fmt.Printf("type of this rating: %T ", input)
}
