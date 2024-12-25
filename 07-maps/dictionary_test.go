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

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "new"
		definition := "new definition"
		dictionary.Add(key, definition)

		assertDefinition(t, dictionary, key, definition)
	})

	t.Run("add repeated word", func(t *testing.T) {
		key := "key"
		definition := "definition"
		dictionary := Dictionary{key: definition}
		err := dictionary.Add(key, "new key definition")

		// if err == nil {
		// 	t.Fatal("expected error but didn't get one")
		// }

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, definition)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, definition string) {
	got, err := dictionary.Search(key)
	want := definition

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, want)
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
