package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "test definition"}
	got := Search(dictionary, "test")
	want := "test definition"

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}
