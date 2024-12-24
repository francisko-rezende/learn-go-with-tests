package structsmethodsandinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(4.0, 2.0)
	want := 12.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// So far, so easy. Now let's create a function called Area(width, height float64) which returns the area of a rectangle.
func TestArea(t *testing.T) {
	got := Area(2.0, 2.0)
	want := 4.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
