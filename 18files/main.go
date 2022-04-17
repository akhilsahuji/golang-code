package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files")
	content := "This is content for my file."

	file, err := os.Create("./myfirstfileingo.txt")
	checkNillErr(err)
	length, err := io.WriteString(file, content)
	checkNillErr(err)
	fmt.Println("The length is: ", length)
	defer file.Close()
	readFile("./myfirstfileingo.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNillErr(err)
	fmt.Println("Data inside the file is: ", string(databyte))
}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
