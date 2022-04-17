package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
)

func main() {
	fmt.Println("Welcome to the maths in golang")

	//var myNumberOne int = 3
	//var myNumberTwo float64 = 4

	//fmt.Println("The Sum is:", myNumberOne+int(myNumberTwo))

	//random number using math/rand

	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(5)+1)

	//random number generator using crypto/rand

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
