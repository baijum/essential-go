package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Errorf("First error and continue")
	t.Fatalf("Second error and not continue")
	t.Errorf("Third error does not display")
}
