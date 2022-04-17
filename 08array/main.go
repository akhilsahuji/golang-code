package main

import "fmt"

func main() {
	fmt.Println("Wellcome to array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [5]string{"potato", "beans", "mushrooms"}
	fmt.Println("Vegy list is: ", len(vegList))
}
