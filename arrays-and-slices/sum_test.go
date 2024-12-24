package arraysandslices

import (
	"slices"
	"testing"
)

func checkSliceSums(t testing.TB, got, want []int) {
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

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

func TestSumAll(t *testing.T) {
	t.Run("Works with two slices", func(t *testing.T) {
		firstSlice := []int{1, 2}
		secondSlice := []int{0, 9}

		got := SumAll(firstSlice, secondSlice)
		want := []int{3, 9}

		checkSliceSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sums the tail of different slices", func(t *testing.T) {
		firstSlice := []int{1, 2, 3}
		secondSlice := []int{0, 9}

		got := SumAllTails(firstSlice, secondSlice)
		want := []int{5, 9}

		checkSliceSums(t, got, want)
	})

	t.Run("works on empty slices", func(t *testing.T) {
		firstSlice := []int{1, 2, 3}
		secondSlice := []int{}

		got := SumAllTails(firstSlice, secondSlice)
		want := []int{5, 0}

		checkSliceSums(t, got, want)
	})
}
