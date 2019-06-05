package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"mykey": "my value"}
	t.Run("known", func(t *testing.T) {
		got, _ := dictionary.Search("mykey")
		want := "my value"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}
func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s', wanted '%s'", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}
