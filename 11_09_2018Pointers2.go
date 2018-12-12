package main

import "fmt"

func main() {
	const max int = 4
	arr := [max] int{}
	var ptr1 [max]*int;
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40

	fmt.Println("ptr1", ptr1)
	//fmt.Println("*ptr1",*ptr1)
	fmt.Println("&ptr1", &ptr1)
	fmt.Println("arr", arr)
	fmt.Println("&arr[0]", &arr[0])
	fmt.Println("-------------after---------")
	var index int
	for index = 0; index < max; index++ {
		fmt.Println("--------------------------")
		ptr1[index] = &arr[index] /* assign the address of integer. */
		fmt.Println("ptr1[", index, "]", ptr1[index])
		fmt.Println("*ptr1[", index, "]", *ptr1[index])
		fmt.Println("&ptr1[", index, "]", &ptr1[index])
		fmt.Println("arr[", index, "]", arr[index])
		fmt.Println("&arr[", index, "]", &arr[index])
	}
}
