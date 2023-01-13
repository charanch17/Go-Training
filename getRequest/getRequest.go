package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, _ := http.Get("https://lco.dev/")
	databytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(databytes))
	defer response.Body.Close()
}
