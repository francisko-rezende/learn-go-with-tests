package main

import (
	"os"

	"github.com/francisko-rezende/learn-go-with-tests/mocking"
)

func main() {
	// fmt.Println("Learning go with tests")
	mocking.Countdown(os.Stdout)
}
