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
	lang := language("en") //? #A
	want := "Hello world"  //? #A

	got := greet(lang) //? #B

	if got != want { //? #C
		// mark this test as failed
		t.Errorf("expected %s, got %s", want, got)
	}
}

func TestGreet_French(t *testing.T) {
	lang := language("fr")     //? #A
	want := "Bonjour le monde" //? #A

	got := greet(lang) //? #B

	if got != want { //? #C
		// mark this test as failed
		t.Errorf("expected %s, got %s", want, got)
	}
}

// Akkadian is not implemented yet!
func TestGreet_Akkadian(t *testing.T) {
	lang := language("akk") //? #A
	want := ""              //? #A

	got := greet(lang) //? #B

	if got != want { //? #C
		// mark this test as failed
		t.Errorf("expected %s, got %s", want, got)
	}
}
