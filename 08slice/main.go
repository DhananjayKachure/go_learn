package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("hello this is slice")
	var array = []string{"choco"}
	var fruitList = []string{"apple", "Tomato", "Peach"}
	fruitList = append(fruitList, array...)
	fmt.Println(fruitList) //[apple Tomato Peach choco]
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList) //[Tomato Peach]
	var num = []int{1, 5, 7, 6, 9, 0}
	highScore := make([]int, 0)
	fmt.Println("this is high score", highScore)
	highScore = append(highScore, num...)
	fmt.Println("this is high score", highScore)
	sort.Ints(highScore)
	fmt.Println("this is high score", highScore)

	//how to remove from slices
	var cources = []string{"rect", "nexr", "house", "css", "js"}
	var index int = 1
	cources = append(cources[:index], cources[index+1:]...)
	fmt.Println("this is high cources", cources)
}
