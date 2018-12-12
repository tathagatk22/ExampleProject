package main

import "fmt"

func function(ch chan int) {
	for i := 10; i > 0; i-- {
		ch <- i
	}
	close(ch)
}
func main() {
	ch := make(chan int)
	go function(ch)
	for v := range ch {
		fmt.Println("Received ",v)
	}
	//           also
	//for {
	//	v, ok := <-ch
	//	if ok == false {
	//		break
	//	}
	//	fmt.Println("Received ", v, ok)
	//}
}
