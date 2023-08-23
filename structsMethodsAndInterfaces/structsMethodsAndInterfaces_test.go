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
		shape Shape
		want  float64
	}{
		{Rectangle{5.0, 2.0}, 10.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
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
