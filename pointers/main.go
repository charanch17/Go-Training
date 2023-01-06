package main

import "fmt"

func add(x, y *int) int {
	fmt.Println("before", *x, *y)
	*x = *x + 5
	return *x + *y

}

func main() {
	x := 2
	y := 1
	fmt.Println(add(&x, &y))
	fmt.Println("after", x, y)
}
