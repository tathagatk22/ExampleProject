package main

import ("math"
		"fmt")

/* define a circle */
type Circle struct {
	x, y, radius float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	var bal =[] int{1222,43434}
	fmt.Println(len(bal))
	fmt.Printf("Circle area: %f", circle.area())
}
