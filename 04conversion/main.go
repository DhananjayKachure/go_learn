package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a)
	fmt.Println("welcome to our pizza app")
	fmt.Println("rae between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating, ", input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("herer is the raing", numRating+1)
	}

}
