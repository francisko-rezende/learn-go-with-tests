package structsmethodsandinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{4.0, 2.0}
	got := Perimeter(rectangle)
	want := 12.0

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{2.0, 2.0}, 4.0},
		{Circle{10.0}, 314.1592653589793},
		{Triangle{2.0, 5.0}, 5.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
