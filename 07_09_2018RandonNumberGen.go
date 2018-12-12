package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // you only need to set seed one time , to generate different values for each and every time when randint is called
	i := 0
	for i <= 3 {
		fmt.Println(rand.Int())
		i++
	}
}
