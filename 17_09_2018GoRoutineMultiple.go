package main

import (
	"fmt"
	"time"
)

func Numbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(2 * time.Millisecond)
		fmt.Println("Numbers", i)
	}
}
func Alphabets() {
	for i := 1; i <= 10; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println("Alphabets", i)
	}
}
func main() {
	go Numbers()
	go Alphabets()
	for i := 1; i <= 5; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("main ", i)
	}

}
