package main

import "fmt"

const LoginToken string = "lkdldf" //if variable first letter is capital it will be declare as public variable
func main() {
	var username string = "Dhananjay"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallNumber uint8 = 255
	fmt.Println(smallNumber)
	fmt.Printf("variable is of type: %T \n", smallNumber)

	var number int = 2554545646
	fmt.Println(number)
	fmt.Printf("variable is of type: %T \n", number)

	var smallFloat float32 = 255.4545646
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	// implicit type
	var name = "Dhananjay"
	fmt.Println(name)
	fmt.Printf("variable is of type: %T \n", name)

	//no var style

	numberOfUser := 123456 // := wallerous opreator  its not allowed outside a module
	fmt.Println(numberOfUser)
	fmt.Printf("variable is of type: %T \n", numberOfUser)

	//const
	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)

}
