package main

import "fmt"

func function(value string) {
	for i := 0; i < 3; i++ {
		fmt.Println("i = ", i, " value = ", value)
	}
}
func main() {

	function("this is direct")

	go function("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going in anonymous function  1")

	//fmt.Scanln()
	fmt.Println("done")
}
