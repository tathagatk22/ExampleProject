package main

import "math"
import "fmt"

func main() {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))
	getValue := func(abc int) {
		fmt.Println(abc + abc)

	}
	getValue(123)
	getAge := func() {
		fmt.Print("Tatghagafrt")
	}
	getAge()
}
