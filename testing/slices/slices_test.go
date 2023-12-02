package slices

import (
	"reflect"
	"testing"
)

// go test -cover, gives you the amount of passed tests
func TestSum(t *testing.T) {
	//funcion como variable, permite mantener el scope dentro de la funcion en la que fue
	//declarada, en este caso no podria usarse en otro lugar que no sea TestSum
	checkSum := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	}

	t.Run("sums", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3, 4, 5, 6})
		expected := 21

		if got != expected {
			t.Errorf("got %q expected %q", got, expected)
		}
	})
	t.Run("sums", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3, 4, 5, 6, 7})
		expected := 28
		if got != expected {
			t.Errorf("got %d expected %d", got, expected)
		}
	})
	t.Run("sums", func(t *testing.T) {
		got := SumMultiple([]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 1, 1}, []int{2})
		expected := []int{28, 3, 2}
		checkSum(t, got, expected)
	})
	t.Run("sums", func(t *testing.T) {
		got := SumMultipleTails([]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 1, 1}, []int{2})
		expected := []int{27, 2, 0}
		checkSum(t, got, expected)
	})
}
