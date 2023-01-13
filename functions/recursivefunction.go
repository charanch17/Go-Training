package main

import "fmt"

var num_list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func sum(i int, total int) int {
	if i == len(num_list) {
		return total
	}
	return sum(i+1, total+num_list[i])

}

func main() {

	fmt.Println(sum(0, 0))

}
