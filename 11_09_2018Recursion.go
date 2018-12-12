package main

import "fmt"

//fibonanci series example
func main() {
	var number int = 8
	fmt.Println(fact(number))
	fmt.Println(fibo(number))
}
//fibonanci series example
func fibo(number int) int{
	if number <= 2{
		return  1
	}
	return fibo(number - 1) + fibo(number - 2)
}
func fact(number int) int  {
	if(number <= 1){
		return 1
	}
	return number *	fact(number - 1)
}