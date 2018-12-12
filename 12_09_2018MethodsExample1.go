package main

import "fmt"

type Square struct {
	x float64
	y float64
}

func exchangeValuesWithoutPointer(s Square) {
	temp := s.x
	s.x = s.y
	s.y = temp
}
func exchangeValuesWithPointer(s *Square) {
	temp := s.x
	s.x = s.y
	s.y = temp
}
func (s Square) exchange() {
	temp := s.x
	s.x = s.y
	s.y = temp
}
func (s *Square) exchangeWithPointer() {
	// There are two reasons to use a pointer receiver.
	// The first is so that the method can modify the value that its receiver points to.
	// The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct.
	temp := s.x
	s.x = s.y
	s.y = temp
}

// https://tour.golang.org/methods/6
func main() {
	square := Square{12.35, 986.3}
	fmt.Println("original ", square)

	exchangeValuesWithoutPointer(square)
	fmt.Println("exchangeValuesWithoutPointer ", square)

	exchangeValuesWithPointer(&square)
	fmt.Println("exchangeValuesWithPointer ", square)

	fmt.Println("---------------------")
	fmt.Println("original ", square)

	square.exchange()
	fmt.Println("exchangeValues ", square)

	square.exchangeWithPointer()
	fmt.Println("exchangeValuesPointer ", square)

	squareNew := &Square{454.454, 9898.23}
	fmt.Println("---------------------")
	fmt.Println("original ", squareNew)

	squareNew.exchange()
	fmt.Println("exchangeValues ", squareNew)

	squareNew.exchangeWithPointer()
	fmt.Println("exchangeValuesPointer ", squareNew)

	fmt.Println("---------------------")
	//exchangeValuesWithoutPointer(squareNew) //(type *Square) as type Square
	//fmt.Println("exchangeValuesWithoutPointer ", squareNew)

	exchangeValuesWithPointer(squareNew)
	fmt.Println("exchangeValuesWithPointer ", squareNew)

}
