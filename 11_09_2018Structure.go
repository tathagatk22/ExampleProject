package main

import "fmt"

type Person struct {
	name  string
	age   int
	phone int64
}

//Reference is needed to set variables
func setPerson(person *Person, name string, phone int64) {
	person.name = name
	person.phone = phone
}
//Reference is not needed, changes into this scope will not reflect to original object
func showPerson(person Person) {
	fmt.Println(person.name)
	fmt.Println(person.age)
	fmt.Println(person.phone)
}
func main() {
	var person1 Person
	var person2 Person
	setPerson(&person1, "Tathagat", 8390972124)
	setPerson(&person2, "Tathagat", 8390972124)
	showPerson(person1)
	showPerson(person2)
}
