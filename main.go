package main

import "fmt"

func main() {
	// fmt.Println("Hello world")
	greeting := greet()
	fmt.Println(greeting)
}

func greet() string {
	return "Hello world"
}
