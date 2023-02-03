package main

import "fmt"

type nicePerson struct {
	name         string
	phoneNumber  string
	emailAddress string
}

func newPerson(name string, phoneNumber string, emailAddress string) {
	person := nicePerson{name, phoneNumber, emailAddress}
	fmt.Println(person.name)
	fmt.Println(person.phoneNumber)
	fmt.Println(person.emailAddress)
	fmt.Println(&person)
	p := &person
	fmt.Println(p)
	fmt.Println("Hello World!")
}

func main() {
	newPerson("Some Person", "990,000,000", "anemail@email.com")
}
