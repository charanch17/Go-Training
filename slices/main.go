package main

import "fmt"

func main() {
	fruitsList := []string{"apple", "mango", "tomato"}
	// adding elements
	fruitsList = append(fruitsList, "carrot", "beetroot")

	fmt.Println(fruitsList)

	// adding at specific loc/index =>2

	fruitsList = append(fruitsList[:2], append([]string{"plum"}, fruitsList[2:]...)...)
	fmt.Println(fruitsList)

	// delete at particular index

	fruitsList = append(fruitsList[:3], fruitsList[4:]...)
	fmt.Println(fruitsList)

}
