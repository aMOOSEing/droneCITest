package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	got := Greeting("Mus")
	want := "Hello, Mus"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
