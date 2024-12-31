package mocking

import (
	"fmt"
	"io"
)

func Countdown(writer io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
	}
	fmt.Fprint(writer, "Go!")
}
