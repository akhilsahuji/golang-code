package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 24

	var result string

	if loginCount > 10 {
		result = "Watch out"

	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println(result)

	if 6%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if num := 11; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}

}
