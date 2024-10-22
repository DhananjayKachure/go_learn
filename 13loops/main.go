package main

import "fmt"

func main() {
	fmt.Println("lets learn Loop")
	days := []string{"sun", "mon", "tue", "wed", "sat"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, value := range days {
		fmt.Printf("index is %v value is %v\n", index, value)
	}

	rougevalue := 1
	for rougevalue < 10 {

		if rougevalue == 2 {
			goto lco

		}
		if rougevalue == 5 {
			rougevalue++
			continue
		}
		fmt.Println("value is", rougevalue)
		rougevalue++

	}

lco:
	fmt.Println("juming to me")

}
