package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// writing to file
	file, _ := os.Create("./test.txt")
	length, _ := io.WriteString(file, "test file")
	fmt.Println(length)
	// reading from a file
	databytes, _ := ioutil.ReadFile("./test.txt")
	fmt.Println(string(databytes))
}
