package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello world
}

// ? A: preparation phase
// ? B: execution phase
// ? C: decision phase
func TestGreet_English(t *testing.T) {

	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{ //? #A
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Spanish": {
			lang: "es",
			want: "Hola mundo",
		},
		"Akkadian, not supported": {
			lang: "akk",
			want: `Unsupported language: "akk"`,
		},
		"Greek": {
			lang: "el",
			want: "Χαίρετε Κόσμε",
		},
		"Hebrew": {
			lang: "he",
			want: "עולם שלום",
		},
		"Urdu": {
			lang: "ur",
			want: "ﺎﯿﻧد ﻮﻠﯿہ",
		},
		"Vietnamese": {
			lang: "vi",
			want: "Xin chào Thế Giới",
		},
		"Empty": {
			lang: "",
			want: `Unsupported language: ""`,
		},
	}

	for name, tc := range tests {

		// Range over all test cases
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang) //? #B
			if got != tc.want {   //? #C
				// mark this test as failed
				t.Errorf("Expected %q, got %q", tc.want, got)
			}

		})

	}

}
