package arraysandslices

import (
	"slices"
	"testing"
)

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

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	firstSlice := []int{1, 2, 3}
	secondSlice := []int{0, 9}

	got := SumAllTails(firstSlice, secondSlice)
	want := []int{5, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
