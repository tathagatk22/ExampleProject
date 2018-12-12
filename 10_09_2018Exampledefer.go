package main

import (
	"fmt"
)

func main() {
	//defer fmt.Println(100 , 100)
	fmt.Println("sasas")
	//defer function1()
	//saveMe()
	fmt.Println(addition(2, 3))
	fmt.Println(count(10))
}
func addition(x int, y int) int {

	//defer fmt.Println(x , y)
	//defer fmt.Println(x+100 , y+100)
	fmt.Println("x = ", x , "y = ", y )
	return x + y
}
func function1(){
	defer func(){
		fmt.Println("ljhFvgdfgdfgdfg")
	}()
	fmt.Println("aasasasas")
	return
}
func saveMe() {
	defer func() {
		fmt.Printf("recovered in defer: \"%s\"\n\n", recover())

		// Go saves the stacktrace when calling defer.
		//debug.PrintStack()
	}()

	// if there was no defer func here, the program will terminate.
	// after the following code, the defer will be executed.
	panic("oops!")
}
func count(i int) (n int) {

	// i = 10
	fmt.Println(n,"1")
	fmt.Println(i,"1")
	defer func(n int) {//try changing n to i 4 combinations
		// n = 20, i = 10
		fmt.Println(n,"3")
		fmt.Println(i,"3")
		n = n + i
		fmt.Println(n,"4")
		fmt.Println(i,"4")
		// n = 20 + 10 = 30
	}(n)//try changing n to i total 4 combinations

	i = i * 2
	n = i
	fmt.Println(n,"2")
	fmt.Println(i,"2")
	// i = 10 * 2 = 20
	// n = 20

	return
}
