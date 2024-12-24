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
		name  string
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 2.0, Heigth: 2.0}, want: 9.0},
		{name: "circle", shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 2.0, Heigth: 5.0}, want: 5.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}
