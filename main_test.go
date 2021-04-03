package main

import "testing"

func TestHello(t *testing.T) {
	got := greet("Adnan")
	want := "Hello Adnan"
	
	if got != want {
		t.Errorf("GOT %q WANT %q", got, want)
	}
}