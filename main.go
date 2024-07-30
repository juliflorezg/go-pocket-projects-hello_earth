package main

import "fmt"

func main() {
	// fmt.Println("Hello world")
	greeting := greet("en")
	fmt.Println(greeting)
}

// language represents the language’s code
type language string

func greet(l language) string {

	switch l {
	case "en":
		return "Hello world"
	case "fr":
		return "Bonjour le monde"
	default:
		return ""
	}

}
