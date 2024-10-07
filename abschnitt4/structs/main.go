package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

// func (p person) updateName(newFirstName string) { // copy by value
func (pointerToPerson *person) updateName(newFirstName string) { // copy by ref
	pointerToPerson.firstName = newFirstName // short for (*pointerToPerson).firstName = ...
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
