package main

import (
	"fmt"
)

func main() {
	fmt.Println("lets start with struct")
	detail := User{"Dhananjay", "dhanajay@dev.com", true, 15}
	fmt.Println(detail)
	fmt.Printf("detail: %+v\n", detail)
	detail.getStatus()
	detail.NewMail()
	fmt.Printf("status:%v\n newMail:%v", detail.Status, detail.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus() {
	fmt.Println("is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "tesat@dev"
	fmt.Println("email of this user is ", u.Email)

}
