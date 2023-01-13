package main

import "fmt"

type Person struct {
	firstname string
	lastname  string
	age       int
}

func (p Person) fullname() string {
	return p.firstname + " " + p.lastname

}

func main() {
	x := Person{
		firstname: "charan",
		lastname:  "ch",
	}
	fmt.Println(x.fullname())

}
