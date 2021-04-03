package main

import (
	"fmt"
)

func Greet(name string, lang string) string {
	greetPrefix := getGreetPrefix(lang)
	if name == "" {
		name = "World"
	}

	return greetPrefix + name
}

func getGreetPrefix(lang string) (prefix string) {
	switch lang {
	case "Spanish":
		prefix = "Hola "
	case "French":
		prefix = "Bonjour "
	default:
		prefix = "Hello "
	}
	return
}

// Add takes two integers and returns the sum of them
func Add(a, b int) int {
	return a + b
}

// returns s which contains c n times
func Repeat(c byte, n int) (s string) {
	for i := 0; i < n; i++ {
		s += string(c)
	}

	return
}

// add numbers inside slice
func Sum(s []int) (res int) {
	for _, num := range s {
		res += num
	}

	return
}

// add numbers inside arr
func SumAll(slices ...[]int) (sumSlice []int) {
	for _, slice := range slices {
		sumSlice = append(sumSlice, Sum(slice))
	}
	return
}

func SumAllTails(slices ...[]int) (sumSlice []int) {
	for _, slice := range slices {
		if len(slice) == 1 {
			sumSlice = append(sumSlice, slice[0])
		} else if len(slice) == 0 {
			sumSlice = append(sumSlice, 0)
		} else {
			sumSlice = append(sumSlice, Sum(slice[1:]))
		}
	}

	return
}

func main() {
	fmt.Println(Greet("Adnan", "Spanish"))
}
