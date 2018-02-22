package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   string
}

func main() {
	// alex := person{"Alex", "Anderson"}

	//	alex := person{firstName: "Alex", lastName: "Anderson"}

	//	fmt.Println(alex.firstName, alex.lastName)

	var nueva person
	nueva.firstName = "nombre"
	nueva.lastName = "lastname"
	nueva.contactInfo.email = "email"
	nueva.contactInfo.zip = "zip"

	// nuevaPointer := &nueva
	nueva.updateName("perro")
	fmt.Println(nueva)
	// fmt.Printf("%+v", nueva)

}

func (nuevaPointer *person) updateName(newFirstName string) {
	(nuevaPointer).lastName = newFirstName
}

func (p person) print() {
	fmt.Println("%+v", p)

}
