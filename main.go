package main

import (
	"flag"
	"fmt"
)

func main() {

	var lang string
	flag.StringVar(&lang, "lang", "en", "Value for the greeting language, e.g. en, ru, fr, es ...")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// language represents the language’s code
type language string

var phrasebook = map[language]string{
	// two-letter codes for languages are standardized by ISO 639-1.
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello world",       // English
	"es": "Hola mundo",        // Spanish
	"fr": "Bonjour le monde",  // French
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
