package main

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	age int
	name string
}

func ass(in interface{}){
	s, err := in.(int)
	fmt.Println(in,"   ",err)
	fmt.Println(s,"   ",err)
}
func TypeSwitch(Interface interface{})  {
	switch v := Interface.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case rune:
		fmt.Println("rune")
	case Person:
		v.Describe()
	}
}
func (p Person)Describe()  {
	fmt.Println(p.name ,"   ",p.age)
}
func main() {
	var name interface{}= 12
	ass(name)
	TypeSwitch( 'a')
	TypeSwitch( 32223)
	TypeSwitch( 32223.332321)
	person := Person{23, "tathagat"}
	TypeSwitch(person)
}
