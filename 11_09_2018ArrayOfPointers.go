package main

import "fmt"



func main() {
	const max int = 4
	arr := [max] int{}
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	var ptr *[max]int //or var ptr [max]*int
	fmt.Println("ptr", ptr)
	//fmt.Println("*ptr",*ptr)
	fmt.Println("&ptr", &ptr)
	fmt.Println("arr", arr)
	fmt.Println("&arr[0]", &arr[0])
	fmt.Println("----------------------")
	ptr = &arr
	fmt.Println("ptr",ptr)
	fmt.Println("*ptr",*ptr)
	fmt.Println("&ptr",&ptr)
	fmt.Println("arr",arr)
	fmt.Println("&arr",&arr)
	fmt.Println("-------------------")
	var index int
	fmt.Println("--------------------------")
	for index = 0; index < max ;index++  {
		fmt.Println("--------------------------")
		//fmt.Println("*ptr",*ptr)
		fmt.Println("ptr[",index,"]",ptr[index])
		fmt.Println("*ptr[index]",*ptr)// error
		fmt.Println("&ptr[",index,"]",&ptr[index])
		fmt.Println("arr[",index,"]",arr[index])
		fmt.Println("&arr[",index,"]",&arr[index])
		fmt.Println("--------------------------")
	}

	fmt.Println("--------------------------")
	//for index = 0; index < max; index++ { //ptr[index] is already in use by another variable
	//	ptr[index] = &arr[index] // error : Cannot use '&arr[index]' (type *int) as type int in assignment
	//}
	//a := []int{10,100,200}
	var ptr1 [max]*int;
	for  index = 0; index < max; index++ {
		fmt.Println("--------------------------")
		ptr1[index] = &arr[index] /* assign the address of integer. */
		fmt.Println("ptr1[",index,"]",ptr1[index])
		fmt.Println("*ptr1[",index,"]",*ptr1[index])
		fmt.Println("&ptr1[",index,"]",&ptr1[index])
		fmt.Println("arr[",index,"]",arr[index])
		fmt.Println("&arr[",index,"]",&arr[index])
		//fmt.Println("*arr[",index,"]",*arr[index])//Invalid indirect of 'arr[index]' (type 'int')
	}
}

