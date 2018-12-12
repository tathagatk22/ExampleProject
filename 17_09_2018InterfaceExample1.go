package main

import (
	"fmt"
)

//interface definition
type VowelsFinder interface {
	FindVowels() []rune
}
type Vowel interface {
	Find() string
}
type MyString string

type Value int
//MyString implements VowelsFinder

func (ms Value) Find() string {

	fmt.Println("%%%%%%%%%%%")
	return "saszasasas"
}
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := MyString("Sam Anderson")
	var v VowelsFinder
	fmt.Printf("name =  %T  v= %T \n",name,v)
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c \n", v.FindVowels())
	fmt.Printf("name =  %T  v= %T \n",name,v)

	Value := Value(1212)
	var ar Vowel
	ar = Value
	fmt.Println( ar.Find)
	fmt.Println( ar.Find())
}