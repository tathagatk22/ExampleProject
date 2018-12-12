package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Enter Any number")
	var str string
	fmt.Scanf("%s", &str)
	value ,_:= strconv.ParseInt(str ,10 ,32)
	fmt.Printf("%T\n" , value)
	newValue,_ := strconv.Atoi(str)
	fmt.Printf("%T" , newValue)
	float, _ := strconv.ParseFloat(str, 32)
	fmt.Println(float)
	}
