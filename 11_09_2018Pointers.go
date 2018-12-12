package main

import "fmt"

func main() {
	//var a int = 10
	//fmt.Println(&a)
	var abc int = 20;
	var abcptr *int = &abc
	//abcptr = &abc
	fmt.Println("abc", abc)
	fmt.Println("&abc", &abc)
	fmt.Println("abcptr", abcptr)
	fmt.Println("*abcptr", *abcptr)
	fmt.Println("&abcptr", &abcptr)
	fmt.Print( "%x abcptr  ")
	fmt.Printf("%x \n" , *abcptr)

	fmt.Println("Nil Pointer")

	var ptr *int
	fmt.Println("ptr", ptr)
	fmt.Println("&ptr", &ptr)
	//fmt.Println("*ptr" , *ptr)
	fmt.Print("%x , ptr  ")
	fmt.Printf("%x ", ptr)
	if ptr == nil{
		fmt.Print("Inside If")
	}
}
