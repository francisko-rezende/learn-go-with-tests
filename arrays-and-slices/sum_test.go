package arraysandslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("A collection with any amount of numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
}
