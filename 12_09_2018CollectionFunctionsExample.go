package main

import (
	"strings"
	"fmt"
	)

func Any(vs []string, f func(string) bool) bool {
	fmt.Println("1")
	for _, v := range vs {
		if f(v) {
			fmt.Println("3")
			return true
		}
	}
	fmt.Println("3")
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	fmt.Println("1**********")
	vsm := make([]string, len(vs))
	for i, v := range vs {
		fmt.Println("2**********")
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	strs := []string{"peach", "apple", "pear", "plum"}
	//anonymous functions
	fmt.Println(Any(strs, func(v string) bool {
		fmt.Println("2")
		return strings.HasPrefix(v, "a")
	}))
	//anonymous functions
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	//anonymous functions
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	//anonymous functions
	//anonymous functions
	fmt.Println(Map(strs, func(v string) string {
		fmt.Println("3**********")
		return strings.ToLower(v)
	}))

	fmt.Println(Map(strs, strings.ToUpper))

}
