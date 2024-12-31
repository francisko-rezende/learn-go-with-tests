package main

import (
	mocking "learn-go-with-tests/09-mocking"
	"os"
	"time"
)

func main() {
	defaultSleeper := &mocking.ConfigurableSleeper{Duration: time.Second * 1, SleepFunc: time.Sleep}
	mocking.Countdown(os.Stdout, defaultSleeper)
}
