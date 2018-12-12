package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
func main() {
	ch := make(chan int, 1)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v,"from ch")
		time.Sleep(1 * time.Second)

	}
}
//successfully wrote 0 to ch   //this will come first
//successfully wrote 1 to ch   //this will come first
//read value 0 from ch  	   // this will take time
//successfully wrote 2 to ch
//successfully wrote 3 to ch
//read value 1 from ch
//successfully wrote 4 to ch
//read value 2 from ch
//read value 3 from ch
//read value 4 from ch