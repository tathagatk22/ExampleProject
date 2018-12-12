package main

import (
	"fmt"
	"time"
)

func hello(done chan bool)  {
	fmt.Println("in hello")
	time.Sleep(10 * time.Second)
	fmt.Println("after sleep in hello")
	done <- true
}
//func hi(done chan bool)  {
//	fmt.Println("in hi")
//	time.Sleep(10 * time.Second)
//	fmt.Println("after sleep in hi")
//	done <- true
//}
func main() {

	done:= make(chan bool)
	fmt.Println("to hello")
	go hello(done)
	//go hi(done)
	//fmt.Println("in main")
	//time.Sleep(10 * time.Second)
	//fmt.Println("after sleep in main")
	<- done
	fmt.Println("end in main")
}
