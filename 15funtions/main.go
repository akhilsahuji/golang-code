package main

import "fmt"

func main() {
	fmt.Println("this is funtion")
	greet()

	result := adder(2, 6)

	fmt.Println("Sum is:", result)

	proRes := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The result is: ", proRes)
}

func greet() {
	fmt.Println("Namaste Golang")
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func adder(valOne int, valTwo int) int {

	return valOne + valTwo

}
