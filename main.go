package main

import "fmt"

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

func main() {
	fmt.Println(Greet("Adnan", "Spanish"))
}
