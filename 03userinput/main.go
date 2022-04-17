package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to User Input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating of Pizza:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating, ", input)
	fmt.Printf("Type of rating is %T ", input)
}
