package main

import "fmt"

type User struct {
	UserName   string
	Email      string
	Isloggedin bool
	isActive   bool
}

func main() {

	charan := User{"charanch", "charanch@gmail.com", true, true}
	fmt.Println(charan)
	// gives keys and values
	charan.getStatus()
}

func (u User) getStatus() {
	fmt.Println(u.isActive)

}
