package main

import "testing"

// checks if got == want
func assertCorrectMessage(t testing.TB, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("GOT %#v WANT %#v", got, want)
	}
}

func TestGreet(t *testing.T) {
	t.Run("Default English greeting", func(t *testing.T) {
		got := Greet("Adnan", "Unknown Language")
		want := "Hello Adnan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string name", func(t *testing.T) {
		got := Greet("", "Latin")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Spanish greeting", func(t *testing.T) {
		got := Greet("Adnan", "Spanish")
		want := "Hola Adnan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("French greeting", func(t *testing.T) {
		got := Greet("Adnan", "French")
		want := "Bonjour Adnan"

		assertCorrectMessage(t, got, want)
	})
}
