package main

import (
	mocking "learn-go-with-tests/09-mocking"
	"os"
	"time"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	defaultSleeper := &DefaultSleeper{}
	mocking.Countdown(os.Stdout, defaultSleeper)
}
