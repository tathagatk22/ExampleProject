package main

import "fmt"

//func main() {
//	var a chan int = make(chan int)
//	fmt.Printf("%T ", a)
//}
func hello(done chan bool) {
	for i := 0; i < 5; i++ {
		//done <- true
		fmt.Println("hello")
	}
	done <- true
}
func main() {
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("end of main")

}
