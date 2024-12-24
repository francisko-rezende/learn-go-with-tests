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
		{shape: Rectangle{Width: 2.0, Heigth: 2.0}, want: 4.0},
		{shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{shape: Triangle{Base: 2.0, Heigth: 5.0}, want: 5.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
