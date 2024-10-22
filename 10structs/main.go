package main

import (
	"fmt"
)

func main() {
	fmt.Println("lets start with struct")
	detail := User{"Dhananjay", "dhanajay@dev.com", true, 15}
	fmt.Println(detail)
	fmt.Printf("detail: %+v\n", detail)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
