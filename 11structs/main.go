package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//no inheritance in golang; No super or parent

	akhil := User{"Akhil", "akhil@go.dev", true, 16}
	fmt.Println(akhil)
	fmt.Printf("akhil details are: %+v\n", akhil)
	fmt.Printf("Name is %v and email is %v.", akhil.Name, akhil.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
