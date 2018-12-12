package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) {
	var wordMap = make(map[string]int)
	stringArray := strings.Fields(s)
	for _, value := range stringArray {
		wordMap[value] = len(value)
	}
	fmt.Println(wordMap)
}
func main() {
	var numberMap map[int]string = make(map[int]string)
	numberMap[1] = "one"
	numberMap[2] = "two"
	//fmt.Println(numberMap)
	//numberMap = make(map[int]string)
	//fmt.Println(numberMap)
	for intKey, stringValue := range numberMap {
		fmt.Println("key", intKey, "value", stringValue)
	}
	val, status := numberMap[1]
	if (status) {
		fmt.Println(val)
	} else {
		fmt.Println(status, val)
	}
	for _, stringValue := range numberMap {
		fmt.Println("value", stringValue)
	}
	for intKey := range numberMap {
		fmt.Println("key", intKey, "value", numberMap[intKey])
	}
	delete(numberMap, 3)
	delete(numberMap, 2)
	for intKey := range numberMap {
		fmt.Println("@@@@@@key", intKey, "value", numberMap[intKey])
	}
	WordCount(string("asass hjh hdhjsd    hfhdfhdjf fdf djfhdf f fhd jhfjdhf jf  "))
}
