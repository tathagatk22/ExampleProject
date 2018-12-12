package main

import (
	"fmt"
)

type Vertex struct {
	x float64
	y float64
}

func (v Vertex) getArea() (float64) {
	total := v.y*v.y + v.x*v.x
	return total
}

func getPerimeter(vertex Vertex) float64 {
	total := (vertex.x + vertex.x) * (vertex.y + vertex.y)
	return total
}

func main() {
	vertex1 := Vertex{12.3, 45.6}
	total := vertex1.getArea()
	fmt.Println(total)

	fmt.Println("---------------------")
	vertex32 := &Vertex{12121.989, 78787.65}
	total = vertex32.getArea()
	fmt.Println(total)

	fmt.Println("---------------------")
	total = getPerimeter(*vertex32)
	fmt.Println(vertex32)

	fmt.Println("---------------------")

	vertex32Copy := vertex32
	fmt.Println(vertex32)
	fmt.Println(*vertex32Copy)

	fmt.Println("---------------------")
	vertex32Copy = vertex32
	fmt.Println(*vertex32)
	fmt.Println(vertex32Copy)

	fmt.Println("---------------------")
	vertex32Copy = vertex32
	fmt.Println(*vertex32)
	fmt.Println(*vertex32Copy)

	fmt.Println("---------------------")
	vertex32Copy = vertex32
	fmt.Println(vertex32)
	fmt.Println(vertex32Copy)
	fmt.Println("---------------------")
	vertex32Copy = vertex32
	fmt.Println(*vertex32)
	fmt.Println(*vertex32Copy)

}
