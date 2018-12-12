package main

import "fmt"

type Describe interface {
	describer()
}

func describeInterface(i interface{}) {
	fmt.Printf("\n Type = %T  Value = %v ", i, i)
}
func main() {
	str := "help"
	describeInterface(str)
	i := 55
	describeInterface(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describeInterface(strt)
	var d1 Describe
	if d1 == nil {
		fmt.Printf("\nd1 is nil and has type %T value %v\n", d1, d1)
	}
}
