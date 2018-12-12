package PackageExample

import "math"
import "fmt"
func init(){
	fmt.Println("This is Rectangle Go file")
}
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}