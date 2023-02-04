package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func main() {
	rex := person{
		firstname: "Emmanuel",
		lastname:  "Lubwama",
		contact: contactInfo{
			email:   "lubwamaemmanuel1@gmail.com",
			zipcode: 97966,
		},
	}

	// rexPointer := &rex
	// rexPointer.update("Rex")
	rex.update("Rex")
	rex.print()
}

func (pointerToPerson *person) update(newName string) {
	(*pointerToPerson).firstname = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
