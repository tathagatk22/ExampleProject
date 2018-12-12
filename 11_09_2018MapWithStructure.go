package main

import (
	"fmt"
)

type Position struct {
	long float64
	lat  float64
}

func main() {
	var locationMap map[string]Position = make(map[string]Position)
	locationMap["Nagpur"] = Position{12.21, 32.323}
	locationMap["Mumbai"] = Position{1212.21, 212.3}
	fmt.Println(locationMap)
	//s := string("asass hjh hdhjsd")

	//stringArray := strings.Fields(s)
	//fmt.Println(stringArray)
}
