package main

import "fmt"

const englishPrefix = "Hello, "
const italianPrefix = "Bonjourno, "
const frenchPrefix = "Bonjour, "

// Hello ...
func Hello(name, language string) string {
	var prefix string

	switch language {
	case "English":
		prefix = englishPrefix
		if name == "" {
			name = "World"
		}
	case "Italian":
		prefix = italianPrefix
		if name == "" {
			name = "Mondo"
		}
	case "French":
		prefix = frenchPrefix
		if name == "" {
			name = "Monde"
		}

	}

	return fmt.Sprintf("%v%v", prefix, name)
}

func main() {
	fmt.Println(Hello("world", "English"))
}
