package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var UrlString = "https://lco.dev:3000/learn?coursename=reactjs&userid=1edshui"

func main() {
	result, _ := url.Parse(UrlString)
	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.Query())
	fmt.Println(result.RawQuery)

	urlparts := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "learn",
		RawQuery: "course=react&hello=vfdsk",
	}
	fmt.Println(urlparts.String())

	getUrl := "https://lco.dev/"

	result1, _ := http.Get(getUrl)

	content, _ := ioutil.ReadAll(result1.Body)
	fmt.Println(string(content))

}
