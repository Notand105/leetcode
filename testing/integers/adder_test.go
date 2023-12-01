package integers

import (
	"reflect"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("simple sum", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		if sum != expected {
			t.Errorf("expected %d, got %d ", sum, expected)
		}
	})
	t.Run("double sum", func(t *testing.T) {
		sum := DoubleAdd(2, 2)
		expected := []int{4, -4}
		//para comparar slices podemos hacerlo con reflext.deepequal
		if !reflect.DeepEqual(sum, expected) {
			t.Errorf("expected %d, got %d ", sum, expected)
		}
	})
}
