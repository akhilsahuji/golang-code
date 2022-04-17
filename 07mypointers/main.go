package main

import "fmt"

func main() {

	fmt.Println("Welcome to Pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is: ", ptr)

	myNumber := 23
	//"&" is a referncre sign
	var ptr = &myNumber

	fmt.Println("Value of  Pointer is", ptr)
	fmt.Println("Value of  Pointer is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is:", myNumber)

}
