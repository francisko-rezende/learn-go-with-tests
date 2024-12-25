package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "test definition"}
	got := Search(dictionary, "test")
	want := "test definition"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
