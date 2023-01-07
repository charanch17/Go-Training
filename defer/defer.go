package main

import "fmt"

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%v", i)
	}
}
func main() {
	// follows lifo
	defer fmt.Println("hello")
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("heyyy")
	defer mydefer()

}
