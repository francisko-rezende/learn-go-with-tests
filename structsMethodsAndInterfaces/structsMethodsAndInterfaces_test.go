package structsmethodsandinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("got %g but expected %g", got, expected) //The f is for our float64 and the .2 means print 2 decimal places.
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 5.0, Height: 2.0}, hasArea: 10.0, name: "Rectangle"},
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793, name: "Circle"},
		{shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0, name: "Triangle"},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v, got %g want %g", tt, got, tt.hasArea)
			}
		})
	}
}
