package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	sendch := make(chan int)    //make(chan<- int) gives error
	go sendData(sendch)
	fmt.Println(<-sendch)
}