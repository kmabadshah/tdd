package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("GOT %q WANT %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Greet("Adnan", "English")
		want := "Hello Adnan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string argument", func(t *testing.T) {
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

	t.Run("Default English greeting", func(t *testing.T) {
		got := Greet("Adnan", "Unknown Language")
		want := "Hello Adnan"

		assertCorrectMessage(t, got, want)
	})
}
