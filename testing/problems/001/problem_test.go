package main

import (
	"math"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		got := Perimeter("c", 2)
		want := 2 * math.Pi * 2
		if got != want {
			t.Errorf("got : %f != want: %f", got, want)
		}
	})
	t.Run("case 2", func(t *testing.T) {
		got := Perimeter("s", 2)
		want := 1.0 * 2 * 2
		if got != want {
			t.Errorf("got : %f != want: %f", got, want)
		}
	})
	t.Run("case 3", func(t *testing.T) {
		got := Perimeter("c", 2.5)
		want := 2 * math.Pi * 2.5
		if got != want {
			t.Errorf("got : %f != want: %f", got, want)
		}
	})
}
