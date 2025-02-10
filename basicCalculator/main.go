package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// var calc string = "(1+(2*3))"
	// calcarr := []rune(calc)

	// var openbracs, closebracs []int // Slices to store positions

	// for i := 0; i < len(calcarr); i++ {
	// 	if calcarr[i] == '(' { // Store all '(' positions
	// 		openbracs = append(openbracs, i)
	// 	}

	// 	if calcarr[i] == ')' { // Store all ')' positions
	// 		closebracs = append(closebracs, i)
	// 	}
	// 	for i := 0; i < count; i++ {

	// 	}
	// }

	// fmt.Println("All '(' positions:", openbracs)
	// fmt.Println("All ')' positions:", closebracs)

	user := User{Name: "dh", Age: 28}
	log.Println(user)
	json, _ := json.Marshal(user)
	log.Println(json, "kdm")
	resp, _ := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(json))

	defer resp.Body.Close()
	// body, _ := io.ReadAll(resp.Body)
	log.Print(resp.Status)
}
