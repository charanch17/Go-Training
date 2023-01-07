package main

import "fmt"

func main() {
	mapEx := make(map[string]string)
	mapEx["py"] = "python"
	mapEx["rb"] = "ruby"
	mapEx["js"] = "javascript"
	fmt.Println(mapEx)

	delete(mapEx, "rb")
	fmt.Println(mapEx)
}
