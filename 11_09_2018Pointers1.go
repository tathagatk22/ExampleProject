package main

import "fmt"


func main() {
	const max int = 4
	arr := [max] int{}
	var ptr *[max]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	fmt.Println("ptr", ptr)
	//fmt.Println("*ptr",*ptr)
	fmt.Println("&ptr", &ptr)
	fmt.Println("arr", arr)
	fmt.Println("&arr[0]", &arr[0])
	fmt.Println("-------------after---------")
	ptr = &arr
	fmt.Println("ptr",ptr)
	fmt.Println("*ptr",*ptr)
	fmt.Println("&ptr",&ptr)
	fmt.Println("arr",arr)
	fmt.Println("&arr",&arr)
	var index int
	for index = 0; index < max ;index++  {
		fmt.Println("--------------------------")
		fmt.Println("ptr[",index,"]",ptr[index])
		//fmt.Println("*ptr[index]",*ptr[index])// error
		fmt.Println("&ptr[",index,"]",&ptr[index])
		fmt.Println("arr[",index,"]",arr[index])
		fmt.Println("&arr[",index,"]",&arr[index])
	}
}
