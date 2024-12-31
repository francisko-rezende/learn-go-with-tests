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
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpyCountdownOperations{}

		Countdown(buffer, spySleeper)
		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("sleep before each print", func(t *testing.T) {
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
