package main

import "fmt"

func main() {
	// fmt.Println("Hello world")
	greeting := greet("en")
	fmt.Println(greeting)
}

// language represents the language’s code
type language string

var phrasebook = map[language]string{
	// two-letter codes for languages are standardized by ISO 639-1.
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"es": "Hola mundo",        // Spanish
	"he": "עולם שלום",         // Hebrew
	"ur": "ﺎﯿﻧد ﻮﻠﯿہ",         // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

func greet(l language) string {
	greeting, ok := phrasebook[l]

	if !ok {
		return fmt.Sprintf("Unsupported language: %q", l)
	}

	return greeting
}
