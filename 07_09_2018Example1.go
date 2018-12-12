package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

var c, python, java bool

const constValue int = 50330;

func main() {
	//var points []int  // uninitialized slice

	name := `Tathagat
    abcbdfbhdf`
	fmt.Println(name)
	fmt.Println(len(name))
	age := 4
	fmt.Println(age)
	age = 5
	fmt.Println(age)
	b := 90
	//fmt.Println("changed b is", b, "c is", c)
	b = 32132 * 4245
	fmt.Println(b + 1)
	var d float64 = math.Pow10(10)
	fmt.Println(d + 1)
	var avc string = "assasasasasas" + "wqwqwqwqw";
	fmt.Println(avc)
	fmt.Println("assasasasasas"+"wqwqwqwqw", 2)
	fmt.Println(d)
	var steps [3]string = [3]string{"SEND", "RCVD"}
	fmt.Println(len(steps))
	value := []string{"SEND", "RCVD", "WAIT"}
	//println(value, value, value)
	//print(value, value, value)
	//print("\n")
	//println(value)
	//println(age)
	fmt.Printf("gsdgsdgfhgfhjdsgfhj")
	fmt.Println(value)
	points := make([]int, 4);
	points[0] = 12
	fmt.Println(points)
	const constValue = 500;
	fmt.Println(constValue)
	fmt.Println()
	//sscanf
	var variable float64 = 7878781.145697987
	fmt.Printf("value  %4.2f", variable)
	value1 := "Tathagat";
	value2 := "Khanorkar";
	//booleanVal := false
	defer fmt.Println("world")
	defer fmt.Println(add(2,4))
	defer fmt.Println(add(2,5))
	// Use format string to generate string.
	fmt.Printf("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&%t \n",value1)
	result := fmt.Sprintf("I saw %v  %v\n", value1, value2)
	fmt.Print(result)

	stringValue := fmt.Sprintln("Tathagat", "Khanorkar", "\n", "Age = ", 23)
	fmt.Print(stringValue)
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rand.Intn(111))

	fmt.Println(add(42, 13))
	var varia float64 = float64(constValue * 0.35487)
	fmt.Println(varia)
	var i int
	fmt.Println(i, c, python, java)
	fmt.Println(math.Sqrt(float64(varia * constValue)))
	u := uint(30 * 25554)
	fmt.Print(u)
	fmt.Print("abc", "def", "ghi")
	fmt.Println("abc", u)

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	fmt.Println()
	fmt.Println(time.Now().Weekday().String())
	fmt.Print(time.Month(1))
	//fmt.Println(time.Now())
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")

	}
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")
}
func add(x int, y int) int {

	defer fmt.Println("\nInside Function")
	defer fmt.Println(constValue)
	return x + y
}
