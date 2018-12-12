package main

import (
	"fmt"
	"time"
)
func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	fmt.Println(time.Now().Nanosecond())
	go hello()
	time.Sleep(1 )
	fmt.Println("main function")
	fmt.Println(time.Now().Nanosecond())
	fmt.Println(time.Second * 1)

}