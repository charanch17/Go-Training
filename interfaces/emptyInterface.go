package main

import "fmt"

type EmptyInterface interface{}

func TypeChecker(i EmptyInterface) {
	switch i.(type) {
	case int:
		fmt.Println("int type implements empty interface")
	case string:
		fmt.Println("string type implements empty interface")

	case bool:
		fmt.Println("bool type implements empty interface")

	}
}

func main() {
	TypeChecker("heloo")
	TypeChecker(1)
	TypeChecker(false)

}
