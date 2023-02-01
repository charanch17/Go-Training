package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var fibnacciSeries []int
var seriesDict = make(map[int]int)

func isPalindrome(val int) bool {
	valString := strconv.Itoa(val)
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

func fibanacciGenerator(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		val, exists := seriesDict[n]
		if exists {
			return val
		} else {
			val = fibanacciGenerator(n-1) + fibanacciGenerator(n-2)
			seriesDict[n] = val
			return val
		}
	}
}

func main() {
	fmt.Println("enter a int")
	reader := bufio.NewReader(os.Stdin)
	count, _ := reader.ReadString('\n')
	size, _ := strconv.ParseInt(strings.TrimSpace(count), 10, 64)
	for i := 0; i < int(size); i++ {
		fibnacciSeries = append(fibnacciSeries, fibanacciGenerator(i))
	}
	fmt.Println(fibnacciSeries)
	// palindromic fibanocci in the above series
	for i := 0; i < len(fibnacciSeries); i++ {
		if isPalindrome(fibnacciSeries[i]) {
			fmt.Printf("%v ", fibnacciSeries[i])
		}
	}
}
