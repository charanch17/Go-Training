package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString('\n')
	fmt.Println(data, err)

}
