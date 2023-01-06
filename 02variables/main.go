package main

import "fmt"

func main() {
	// var
	var intValue int = 245

	// implicit declaration
	var stringValue = "this is a stringh"

	var floatValue = 234.023

	// just declaration
	var boolValue bool

	// DEFAULTS
	var defInt int
	var defString string
	var defBool bool
	var defFloat float64

	//volrus op
	x := 1

	fmt.Println("%T %T %T %T", intValue, stringValue, boolValue, floatValue)
	fmt.Println(defInt, defBool, defFloat, defString, x)
}
