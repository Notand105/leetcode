package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("perimeter", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		got := rect.Perimeter()
		var want float64 = 40
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("area", func(t *testing.T) {
		rect := Rectangle{10, 10}
		var want float64 = 100
		checkArea(t, rect, want)
	})
	t.Run("area circles", func(t *testing.T) {
		cir := Circle{10}
		var want float64 = 100 * math.Pi
		checkArea(t, cir, want)

	})
}

func TestArea2(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72},
		{Circle{10}, 100 * math.Pi},
		{Triangle{12, 6}, 36},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}
