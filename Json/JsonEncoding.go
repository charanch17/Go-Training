package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int    `json:"price"`
	Platform string
	Password string `json:"-"`
}

func main() {

	courseData := []course{
		{Name: "react", Price: 10, Platform: "udemy", Password: "fdsfsdfds"},
		{Name: "react", Price: 10, Platform: "udemy", Password: "fdsfsdfds"},
		{Name: "react", Price: 10, Password: "fdsfsdfds"},
	}

	unindentedjsonData, err := json.Marshal(courseData)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(unindentedjsonData))

	indentedJsonData, _ := json.MarshalIndent(courseData, "", "	")
	fmt.Println(string(indentedJsonData))

}
