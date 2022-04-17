package main

import "fmt"

const LoginToken string = "gyahgdh" //public
func main() {
	var isLoginIn bool = true
	fmt.Println(isLoginIn)
	fmt.Printf("Variable is type of %T \n", isLoginIn)

	var username string = "Akhil"
	fmt.Println(username)
	fmt.Printf("Variable is type of %T \n", username)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is type of %T \n", smallVal)

	var smallFloat float32 = 255.355554566
	fmt.Println(smallFloat)
	fmt.Printf("Variable is type of %T \n", smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is type of %T \n", anotherVariable)

	var website = "akhil-sahu.co"
	fmt.Println(website)
	fmt.Printf("Variable is type of %T \n", website)

	numberOfUser := 300000.00
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Varible is type of %T \n", LoginToken)

}
