package structsmethodsandinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("got %g but expected %2.f", got, expected) //The f is for our float64 and the .2 means print 2 decimal places.
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{5.0, 2.0}
		got := rectangle.Area()
		expected := 10.0

		if got != expected {
			t.Errorf("got %g but expected %g", got, expected) // g prints more decimal places
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		expected := 314.1592653589793

		if got != expected {
			t.Errorf("got %g but expected %g", got, expected)
		}
	})
}
