package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Francisko")

	got := buffer.String()
	want := "Hello, Francisko"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
