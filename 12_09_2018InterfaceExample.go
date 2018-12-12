package main

import (
	"fmt"
)

type iShape interface {
	Area() float64
}
type Triangle struct {
	height float64
	base   float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (triangle Triangle) Area() float64 {
	total := triangle.base * triangle.height * 0.5
	return total
}

//func (c Vertex) Area() float64 {
//	return c.x * c.y * 3.14
//}
func (rectangle Rectangle) Area() float64 {
	total := rectangle.length * rectangle.width
	return total
}
func getArea(shape iShape) float64 {
	return shape.Area()
}


func main() {
	triangle := Triangle{5, 6}
	rectangle := Rectangle{5, 4}
	fmt.Println(triangle.Area())
	fmt.Println(rectangle.Area())
	fmt.Println("---------")
	fmt.Println(getArea(triangle))
	fmt.Println(getArea(rectangle))
	//vertex := Vertex{2, 4}
	//fmt.Println(getArea(vertex))
}
