package main

import "testing"

//Tests MUST at a minimum include t *testing.T struct
func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() //Helper test function. Description in docs
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}

	}

	//t.Run is how you basically break up test cases using go's test library anyways
	t.Run("in Espanol", func(t *testing.T) {
		got := Hello("Elodie", "Spanish") //Got calls method or field etc that is being manipulated
		want := "Hola, Elodie"            //Want declares what kind of value or return is EXPECTED from tester
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)

	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Marc", "French")
		want := "Bonjour, Marc"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when empty string provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)

	})

}
