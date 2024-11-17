package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jobs")
	want := "hello Jobs"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
