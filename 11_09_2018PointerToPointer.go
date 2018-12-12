package main

import "fmt"
func main() {
	var number int = 212
	var ptr *int
	var ptrToPtr **int
	var ptrToPtrToPtr ******************************int
	fmt.Println("number", number)
	fmt.Println("&number", &number)
	fmt.Println("ptr", ptr)
	fmt.Println("&ptr", &ptr)
	fmt.Println("ptrToPtr", ptrToPtr)
	fmt.Println("&ptrToPtr", &ptrToPtr)
	fmt.Println("ptrToPtrToPtr", ptrToPtrToPtr)
	ptr = &number
	ptrToPtr = &ptr

	//fmt.Println(ptrToPtrToPtr)
	fmt.Println("---------after-----------")
	fmt.Println("number", number)
	fmt.Println("&number", &number)
	fmt.Println("ptr", ptr)
	fmt.Println("*ptr", *ptr)
	fmt.Println("&ptr", &ptr)
	fmt.Println("ptrToPtr", ptrToPtr)
	fmt.Println("*ptrToPtr", *ptrToPtr)
	fmt.Println("&ptrToPtr", &ptrToPtr)
	fmt.Println("ptrToPtrToPtr", ptrToPtrToPtr)
	//fmt.Println("*ptrToPtrToPtr",*ptrToPtrToPtr)
	//	fmt.Println("&ptrToPtrToPtr" , &ptrToPtrToPtr)
	//	fmt.Println("---------------")
	//	fmt.Println("**ptrToPtr",**ptrToPtr)
	//	fmt.Println("***ptrToPtrToPtr",***ptrToPtrToPtr)
	//}
}