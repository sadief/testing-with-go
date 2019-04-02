package main

import "fmt"

// Hello ...
func Hello(name string) string {
	return fmt.Sprintf("Hello, %v", name)
}

func main() {
	fmt.Println(Hello("world"))
}
