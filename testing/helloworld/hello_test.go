package main

import (
	"testing"
)

// the test function should always start with Test, the function takes only one argument t *	testing.T
// we should import the testing library
// we can force a fail with t.fail()
// you can go documentation locally by running godoc -http :8000
// lo de arriba se necesita instalar mediante sudo apt
func TestHello(t *testing.T) {
	//t.run for multiples tests in the same function
	t.Run("say specific hello", func(t *testing.T) {
		got := HelloWorld("notand", "es")
		want := "Hola notand"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say specific hello", func(t *testing.T) {
		got := HelloWorld("notand", "en")
		want := "Hello notand"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say empty hello", func(t *testing.T) {
		got := HelloWorld("", "it")
		want := "Ciao world"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf(" got %q want %q", got, want)
	}
}
