package main

import "fmt"

func incrementor() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	clsr := incrementor()
	fmt.Println(clsr())
	fmt.Println(clsr())
	fmt.Println(clsr())
	// i will be reset to 0 again if we  call the outer function again
	clsr = incrementor()
	fmt.Println(clsr())

}
