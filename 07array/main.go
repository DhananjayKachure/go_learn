package main

import "fmt"

func main() {

	fmt.Println("welcome to array")
	var fruitList [3]string
	fruitList[0] = "apple"
	fruitList[1] = "mango"
	fruitList[2] = "bannana"
	fmt.Println("fruit array is ", fruitList)

	//another way to declare array

	var veglist = [3]string{"carrot", "potato", "mushroom"}
	fmt.Println("fruit array is ", veglist)
	fmt.Println("fruit array length is ", len(veglist))

}
