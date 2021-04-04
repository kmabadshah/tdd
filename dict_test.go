package main

import "testing"

func TestSearch(t *testing.T) {
	dict := Dict{"test": "this is a test"}

	t.Run("Word exists", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is a test"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Word doesn't exist", func(t *testing.T) {
		_, err := dict.Search("something")

		if err == nil {
			t.Fatal("expected an error")
		}

		assertCorrectMessage(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dict := Dict{"key": "this is just a value"}
	word := "test"
	def := "this is a test"

	t.Run("Add word", func(t *testing.T) {
		_ = dict.Add(word, def)
		got, err := dict.Search(word)
		want := def

		if err != nil {
			t.Fatal("Should've added a new word:", err)
		}
		assertCorrectMessage(t, got, want)
	})

	t.Run("Existing word", func(t *testing.T) {
		err := dict.Add(word, def)

		if err != ErrWordExists {
			t.Fatal("Error expected for existing word: ", err)
		}
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	def := "this is a test"
	dict := Dict{word: def}

	t.Run("Update existing word", func(t *testing.T) {
		newDef := "new word definition"
		err := dict.Update(word, newDef)
		got, _ := dict.Search(word)

		if err != nil {
			t.Fatalf("No error should be thrown as word exists")
		}
		if got != newDef {
			assertCorrectMessage(t, got, newDef)
		}
	})

	t.Run("Update word that does not exist", func(t *testing.T) {
		word := "test2"
		newDef := "this is a another test"

		err := dict.Update(word, newDef)
		if err != nil {
			assertCorrectMessage(t, err, ErrNotFound)
		}
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	def := "this is a test"
	dict := Dict{word: def}

	t.Run("Delete existing word", func(t *testing.T) {
		err := dict.Delete(word)
		if err != nil {
			t.Error("expected no errors: ", err)
		}

		_, err = dict.Search(word)
		if err == nil {
			t.Errorf("Expected word to be deleted")
		}
	})

	t.Run("Delete non existing word", func(t *testing.T) {
		word := "test2"
		err := dict.Delete(word)

		if err == nil {
			t.Errorf("Should've thrown error as word doesn't exist")
		}
	})
}
