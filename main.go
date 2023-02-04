package main

import "fmt"

type List[T any] struct {
	head, tail *listNode[T]
}

type listNode[T any] struct {
	next  *listNode[T]
	value T
}

func (list *List[T]) Push(v T) {
	if list.tail == nil {
		list.head = &listNode[T]{value: v}
		list.tail = list.head
	} else {
		list.tail.next = &listNode[T]{value: v}
		list.tail = list.tail.next
	}
}

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

func buildAList() {
	list := List[int]{}
	list.Push(10)
	fmt.Println(list)
}

func main() {
	newPerson("Some Person", "990,000,000", "anemail@email.com")
	buildAList()
}
