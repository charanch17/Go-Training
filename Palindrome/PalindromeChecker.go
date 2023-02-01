package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(valString string) bool {

	if len(valString) == 1 {
		return true
	}
	leftPtr := 0
	rightPtr := len(valString) - 1
	for leftPtr < rightPtr {
		if valString[leftPtr] != valString[rightPtr] {
			return false

		}
		leftPtr += 1
		rightPtr -= 1
	}
	return true
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter string")
	inputString, _ := reader.ReadString('\n')
	inputString = strings.TrimSpace(inputString)
	fmt.Println(isPalindrome(inputString))
}
