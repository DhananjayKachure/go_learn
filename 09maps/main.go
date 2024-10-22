package main

import "fmt"

func main() {
	fmt.Println("Let's start map")

	// Initializing a map
	languages := make(map[string]string)

	// Adding elements to the map
	languages["js"] = "JavaScript"
	languages["py"] = "Python"
	languages["ja"] = "Java"

	// Printing the map
	fmt.Println(languages)
	fmt.Println(languages["js"])
	delete(languages, "py")
	fmt.Println(languages)

	//loops are intresting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
