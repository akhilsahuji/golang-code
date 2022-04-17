package main

import "fmt"

func main() {
	defer fmt.Println("Welcome to golang defer")
	defer fmt.Println("What is defer?")
	defer fmt.Println("Defer is a keyword that is used when we have to skip a certain line to the end.")
	fmt.Println("Hello World")
	myDefer()
}

func myDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
	}

}
