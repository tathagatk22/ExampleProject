package main

import (
	"fmt"
	"math"
	)

type Person struct {
	name  string
	age   int
	phone int
}

// example 1
func (person Person) showPerson() {
	fmt.Println("Name ", person.name, "\tAge ", person.age, "\tPhone ", person.phone)
}

// example 2
func setPerson(name string, age int, phone int) (Person) {
	return Person{name, age, phone}
}

//example 3
func (person *Person) setAge(age int) {
	person.age = age
}
func (person *Person) setName(name string) {//if pointer person *Person is not added then method operates on a copy of the original person not on actual value(reference)
	person.name = name
}
func (person Person) setPhone(phone int) {
	person.phone = phone
}
func main() {
	//example 1
	person1 := Person{"Tathagat", 45, 7894561231}
	person2 := Person{"Ajay", 75, 3214567898}
	person3 := Person{"Khanorkar", 78, 2357889451}
	person1.showPerson()
	person2.showPerson()
	person3.showPerson()
	//example 2
	person4 := setPerson("Tat", 45, 788787)
	person4.showPerson()
	// example 3
	person4.setAge(4666)
	person4.showPerson()
 	var char rune= 'f'
 	fmt.Printf("\n%T",char)
 	totalNew := int64(9223372036854775807)
 	fmt.Println(totalNew)
 	stringTotal:=string("6739986666787659948666753771754907668409286105635143120275902562304")
	fmt.Println(len(stringTotal))

 	}
