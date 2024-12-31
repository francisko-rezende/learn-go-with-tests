package mocking

import (
	"fmt"
	"io"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func Countdown(writer io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
	}
	fmt.Fprint(writer, finalWord)
}
