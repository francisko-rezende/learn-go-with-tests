package arraysAndSlices

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5} // can also use [...]int{1, 2, 3, 4, 5}
		sum := Sum(numbers)
		expected := 15

		if sum != expected {
			t.Errorf("expected '%d' but got '%d' given '%v'", expected, sum, numbers)
		}

	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("wanted '%d' but got '%d', given '%v'", want, got, numbers)
		}
	})
}
