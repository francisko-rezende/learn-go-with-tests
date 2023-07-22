package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jake")
		want := "Hello, Jake"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
