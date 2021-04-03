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

// takes a byte string and returns a string containing that byte 5 times
func Repeat(c byte, n int) (s string) {
	for i := 0; i < n; i++ {
		s += string(c)
	}

	return
}

func main() {
	fmt.Println(Greet("Adnan", "Spanish"))
}
