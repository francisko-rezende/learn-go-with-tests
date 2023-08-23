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

	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }

	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{5.0, 2.0}
	// 	checkArea(t, rectangle, 10.0)

	// got := rectangle.Area()
	// expected := 10.0

	// if got != expected {
	// 	t.Errorf("got %g but expected %g", got, expected) // g prints more decimal places
	// }
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkArea(t, circle, 314.1592653589793)
	// 	// got := circle.Area()
	// 	// expected := 314.1592653589793

	// 	// if got != expected {
	// 	// 	t.Errorf("got %g but expected %g", got, expected)
	// 	// }
	// })
}
