package main

import "fmt"

type customfunc func(int, int) int

func operation(x, y int, f customfunc) int {
	return f(x, y)

}

func main() {
	output := operation(3, 4, func(x, y int) int {
		return x % y
	})
	fmt.Println(output)

}
