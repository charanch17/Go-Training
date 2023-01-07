package main

import "fmt"

func main() {
	var fruits [4]string
	fruits[0] = "mango"
	fruits[3] = "apple"
	// len will always be the specified one
	fmt.Println(len(fruits), fruits)

}
