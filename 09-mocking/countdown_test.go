package mocking

import (
	"bytes"
	"slices"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

const (
	sleep = "sleep"
	write = "write"
)

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("number of calls", func(t *testing.T) {
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
	})

	t.Run("call order", func(t *testing.T) {
		spy := &SpyCountdownOperations{}

		Countdown(spy, spy)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		got := spy.Calls

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
