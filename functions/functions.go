package main

import "fmt"

func greeter(name string) string {
	return "hello " + name
}

func proadder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val

	}
	return total
}
func multireturner(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(multireturner(2, 3))
	fmt.Println(proadder(1, 2, 3, 4, 5, 6))
	fmt.Println(greeter("charan"))

}
