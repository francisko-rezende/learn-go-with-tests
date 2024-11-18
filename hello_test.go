package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("esaying hello to people", func(t *testing.T) {
		got := Hello("Jobs", "")
		want := "hello Jobs"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'hello world' when an empty string is provided", func(t *testing.T) {
		got := Hello("", "")
		want := "hello world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Juan", "Spanish")
		want := "hola Juan"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("go %q, want %q", got, want)
	}
}
