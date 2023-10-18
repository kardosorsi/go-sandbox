package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson", contactInfo{email: "t@mail.co", zipCode:404}}
	// alex.print()

	// alexa := person{firstName: "Alexa", lastName: "Anderson"}
	// alexa.print()

	// zero values ("", 0) instead of null or undefined
	var alexandra person
	// fmt.Println(alexandra)
	// fmt.Printf("%+v", alexandra)

	alexandra.firstName = "Alexandra"
	alexandra.lastName = "Anderson"
	alexandra.contactInfo.email = "test@mail.com"

	// alexandra.print() 

	alexandraPointer:= &alexandra
	alexandraPointer.updateName("Alexácska")
	// alexandra.print()

	// Go is a pass by value language
	// places the whole thing in another container in the RAM
	// runs the function on the copy as receiver

	jim := person{firstName: "Jim", lastName: "Pointer", contactInfo: contactInfo{email: "test@mail.com"}}
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Println("Hello, my name is " + p.firstName + " " + p.lastName)
	fmt.Println("And my email address is " + p.contactInfo.email)
}