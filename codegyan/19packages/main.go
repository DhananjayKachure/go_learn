package main

import (
	"fmt"
	"log"
	"mypackages/auth"
	"mypackages/user"
)

func main() {
	fmt.Println("lets know about packages")
	auth.Auth("test@.com", "secrete")
	userData := user.UserInfo{
		Email: "testtttting@mail",
		Name:  "Dhananjay",
	}
	log.Println(userData.Name)

}
