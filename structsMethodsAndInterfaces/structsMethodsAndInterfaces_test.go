package structsmethodsandinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	expected := 40.0

	if got != expected {
		t.Errorf("got %.2f but expected %2.f", got, expected) //The f is for our float64 and the .2 means print 2 decimal places.
	}
}

func TestArea(t *testing.T) {
	got := Area(5.0, 2.0)
	expected := 10.0

	if got != expected {
		t.Errorf("got %.2f but expected %.2f", got, expected)
	}
}
