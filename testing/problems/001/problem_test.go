package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		got := HelloWorld("notand", "es")
		want := "Hola notand"
		assertCorrectMessage(t, got, want)
	})
	t.Run("case 2", func(t *testing.T) {
		got := HelloWorld("notand", "en")
		want := "Hello notand"
		assertCorrectMessage(t, got, want)
	})
	t.Run("case 3", func(t *testing.T) {
		got := HelloWorld("", "it")
		want := "Ciao world"
		assertCorrectMessage(t, got, want)
	})
}