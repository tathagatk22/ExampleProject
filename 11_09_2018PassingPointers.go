package main

import "fmt"

func main() {
	number := 10
	number2 := 20
	fmt.Println(&number, &number2)
	fmt.Println(number, number2)
	println("-------")
	exchange(&number, &number2)
	fmt.Println(&number, &number2)
	fmt.Println(number, number2)
}
func exchange(number *int, number2 *int) {
	temp := *number
	*number = *number2
	*number2 = temp
}
