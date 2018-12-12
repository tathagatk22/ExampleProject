package main

import "fmt"
type In interface {
	method()
}
type Stru struct {
	value string 
}
func(stru *Stru)method(){
	if stru == nil {
		fmt.Println("<*******nil>")
		return
	}
	fmt.Println(stru.value,"&&&&&&&&&&")
}
func main() {
	var in2 In

	var structVariable *Stru

	 in2 = structVariable
	
	describe(in2)
	in2.method()

	in2 = &Stru{"hello"}
	describe(in2)

	in2.method()
}
func describe(in In) {
	fmt.Printf("\n %T %v *************" , in,in)
}
