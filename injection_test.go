package main

import (
	"bytes"
	"testing"
)

func TestInjGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	InjGreet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		assertCorrectMessage(t, got, want)
	}
}
