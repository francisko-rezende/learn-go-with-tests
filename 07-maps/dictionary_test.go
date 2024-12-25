package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "test definition"}

	t.Run("simple search", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "test definition"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get an error but didn't get one")
		}

		assertError(t, err, want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
