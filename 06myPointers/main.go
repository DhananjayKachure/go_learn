package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on pointers")
	//how to declare pointer
	//  var ptr *int
	//  fmt.Println("vlue of empty pointer is:" , ptr)
	// output nil
	//hoe to refrence pointer

	myNumber := 5
	var ptr = &myNumber
	fmt.Println("value of refrence pointer", ptr)  // give memory location
	fmt.Println("value of refrence pointer", *ptr) // give actual pointer value

	*ptr = *ptr * 2
	fmt.Println("new value ", myNumber) // it change actual value of myNumber it work on actual that not copy of that value

}
