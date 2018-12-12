package main

import "fmt"
func main() {
	num := 55
	sqr := make(chan int)
	cub := make(chan int)
	go calcSqr(num,sqr)
	go calcCub(num,cub)
	cube, square := <-cub, <-sqr //waiting reading on channel
	fmt.Println( "num = ", num , " square = ",square , " Cube = ", cube)
}
func calcCub(num int, cub chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num /= 10
	}
	cub <- sum // writing on channel
}
func calcSqr(num int, sqr chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	sqr <- sum // writing on channel
}
