package main

import(
	"testing"
)

func TestGreeting(t *testing.T) {
	greeting := Greet("World!")
	expected := "Hello World!"
	if expected != greeting {
		t.Errorf("Greet was incorrect, got: %s, want: %s.", greeting, expected)
	}
}
