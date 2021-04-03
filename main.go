package main

import "fmt"

const englishHelloPrefix = "Hello "

func greet(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(greet("Adnan"))
}
