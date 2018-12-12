package main

import "fmt"

type Name interface {
	getName() string
}
type Age interface {
	getAge() int
}
type ShowPerson interface {
	Name
	Age
}

type Person struct {
	name string
	age  int
}

func (p Person) getName() string {
	return p.name
}
func (p Person) getAge() int {
	return p.age
}
func main() {
	person := Person{"Tathagat", 23}
	fmt.Println("Name", person.getName())
	fmt.Println("Age", person.getAge())
	fmt.Printf("%T \n", person)
	fmt.Printf("%T \n", person.getName())
	fmt.Printf("%T \n", person.getAge())

}
