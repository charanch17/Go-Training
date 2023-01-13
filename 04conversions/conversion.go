package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter a number")
	input, _ := reader.ReadString('\n')
	val, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err == nil {
		fmt.Println(val)
	} else {
		fmt.Println("err", err)
	}
	x := int(3.0)
	fmt.Println(x)

	s := "hello"
	bs := []byte(s)
	fmt.Printf("%T", bs)
	fmt.Println(bs)

}
