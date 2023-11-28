package main

import (
	"testing"
)

// Important to remember that test functions must start with Test
// IMPORTANTLY it must start with a capital T
func Test_HelloWorld(t *testing.T) {
	got := HelloWorld()
	if got != "Hello, World!" {
		t.Fatalf("Wanted Hello, World!, got %v", got)
	}
}
