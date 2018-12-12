package main

import (
	"errors"
	"fmt"
)

func sum(numbers ...int) (int,error) {
	if len(numbers) == 0{
		return 0 ,errors.New(" Zero element in array")
	}
	total := 0
	for _,value := range numbers {
		total = total+value
	}
	return total , nil
}

func main() {
	num1,err :=sum(1,2,3,4,5,5,6,78,9,0,21,3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num1)
	}
	num := []int{1,2,3,4,5,4}
	num1 , err =sum(num...)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num1)
	}
	num2 := []int{}
	num1 , err = sum(num2...)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num1)
	}
}
