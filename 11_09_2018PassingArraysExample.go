package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	var average float64 = getAverage(arr)
	fmt.Println(float64(average))
}
func getAverage(arr []int) float64 {
	var sum int = 0;
	var size int = len(arr)
	for _, value := range arr {
		sum = sum + value
	}
	var average  = float64(sum) / float64(size)
	return average
}
