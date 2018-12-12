package main

import (
	"fmt";
)
func main() {
	/* nextNumber is now a function with i as 0 */
	nextNumber := getSequence()

	/* invoke nextNumber to increase i by 1 and return the same */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* create a new sequence and see the result, i is 0 again*/
	nextNumber = getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
}
func getSequence() func() int {

	i := 0
	return func() int {
		i += 1
		return i
	}
}
