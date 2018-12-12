package main

import "fmt"

func main() {
	srcArray := []int{10, 20, 30, 40, 50, 60, 70, 80, 90} //src array is larger than dst array
	dstArray := []int{100, 200, 300, 400}
	fmt.Println(srcArray)
	fmt.Println(dstArray)
	var total int = copy(dstArray, srcArray) //Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).
	fmt.Println(total)
	fmt.Println(srcArray)
	fmt.Println(dstArray)
	srcArray = []int{10, 20} //src array is smaller than dst array
	dstArray = []int{100, 200, 300, 400, 500}
	fmt.Println(srcArray)
	fmt.Println(dstArray)
	total = copy(dstArray, srcArray) //Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).
	fmt.Println(total)
	fmt.Println(srcArray)
	fmt.Println(dstArray)
	str := "Tathagat Ajay Khanorkar" // it will also copy bytes from a string to a slice of bytes.
	fmt.Println(str)
	byteArray := make([]byte, 10)
	total = copy(byteArray, str) //Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).
	fmt.Printf("%s \n", byteArray)
	fmt.Println(byteArray)
	fmt.Println(str)
	fmt.Println(total)



	srcArray = []int{10,20,30,40,50,60}//The source and destination may overlap.
	fmt.Println(srcArray)
	total = copy(srcArray,srcArray[:])//copy() works correctly even if the destination is a slice which shares the same underlying array as the source slice, and the part of the array designated by source and destination has common parts
	fmt.Println(total)
	fmt.Println(srcArray)

	srcArray = []int{10,20,30,40,50,60}
	fmt.Println(srcArray)
	total = copy(srcArray,srcArray[1:])
	fmt.Println(total)
	fmt.Println(srcArray)

	srcArray = []int{10,20,30,40,50,60}
	fmt.Println(srcArray)
	total = copy(srcArray,srcArray[1:6])
	fmt.Println(total)
	fmt.Println(srcArray)
}
