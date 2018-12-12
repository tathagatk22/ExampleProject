package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	//arrNew := arr
	arrNew := append(arr, 11)
	fmt.Println(arrNew)
	fmt.Println(arr)
	arr[4] = 100
	sliceArr := arrNew[2:5]
	fmt.Println(sliceArr)
	sliceArr[2] = 100
	fmt.Println(sliceArr)
	fmt.Println(arrNew)
	sliceNewArr := append(sliceArr, 78,12)
	fmt.Println(sliceNewArr)
	sliceNewArr[1] = 26865
	fmt.Println(arrNew)
	fmt.Println(sliceArr)
	fmt.Println(sliceNewArr)
	//newArr := make([]int, 10) //make built-in function allocates and initializes an object of type slice, map, or chan (only)
	//fmt.Println(cap(newArr), newArr)
	//newArr = append(newArr, 78)
	//fmt.Println(cap(newArr), newArr)
	//newArr = append(newArr, 78)
	//fmt.Println(cap(newArr), newArr)
	//newArr[4] = 100
	//
	//newArr = append(newArr, 78)
	//newArr = append(newArr, 78)
	//newArr = append(newArr, 78)
	//newArr = append(newArr, 78)
	//newArr = append(newArr, 78)
	//newArr = append(newArr, 78)
	//newArr = append(newArr, 78)
	//newArr = append(newArr, 78)
	//newArr = append(newArr, 78)
	//newArr = append(newArr, 78)
	//newArr = append(newArr, 78)
	//fmt.Println(cap(newArr), newArr)
	//fmt.Println()
	//var names []string
	//if names == nil {
	//	fmt.Println(cap(names), names)
	//	names = append(names, "Tathagat", "Ajay", "Khanorkar")
	//	fmt.Println(cap(names), names)
	//
	//}
	//fruits := []string{"Apple", "Banana"}
	//vegetable := []string{"beans", "Brinjal"}
	//food := append(fruits, vegetable...)
	//fmt.Println(fruits)
	//fmt.Println(food)
	//fmt.Println(vegetable)
}
