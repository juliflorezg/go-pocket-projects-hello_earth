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
func TestGreet(t *testing.T) {
	want := "Hello world" //? #A

	got := greet() //? #B

	if got != want { //? #C
		// mark this test as failed
		t.Errorf("expected %s, got %s", want, got)
	}
}
