package mocking

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)
	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("noto enough calls to spySleeper, wanted 3, got %d", spySleeper.Calls)
	}
}
