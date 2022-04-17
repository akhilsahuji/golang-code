package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://daily.dev"

func main() {
	fmt.Println("LCO web requests")

	response, err := http.Get(url)
	checkErrNil(err)
	fmt.Printf("The Response is type of %T\n", response)
	defer response.Body.Close()
	databyte, err := ioutil.ReadAll(response.Body)
	checkErrNil(err)

	content := string(databyte)

	fmt.Println(content)

	file, err := os.Create("./myfirstfileingo.txt")
	checkErrNil(err)
	length, err := io.WriteString(file, content)
	checkErrNil(err)
	fmt.Println("The length is: ", length)
	defer file.Close()
	// readFile("./myfirstfileingo.txt")

}

// func readFile(filename string) {
// 	databyte, err := ioutil.ReadFile(filename)
// 	checkErrNil(err)
// 	fmt.Println("Data inside the file is: ", string(databyte))

// }
func checkErrNil(err error) {
	if err != nil {
		panic(err)
	}

}
