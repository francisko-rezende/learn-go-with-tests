package main

import "fmt"

const (
	englishHelloPrefix = "hello "
	spanishHelloPrefix = "hola "
	frenchHelloPrefix  = "bonjour "
	spanish            = "Spanish"
	french             = "French"
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("jobs", ""))
}
