package main

import "testing"

func TestWorld(t *testing.T) {
	got := Hello("Jake")
	want := "Hello, Jake"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
