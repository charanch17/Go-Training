package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	urlParts := url.URL{
		Scheme: "https",
		Host:   "jsonplaceholder.typicode.com",
		Path:   "/todos/1",
	}
	resp, err := http.Get(urlParts.String())
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	isValidJson := json.Valid(content)
	if !isValidJson {
		panic("in valide json")

	}
	var jsonData interface{}
	json.Unmarshal(content, &jsonData)

	fmt.Println(jsonData)

}
