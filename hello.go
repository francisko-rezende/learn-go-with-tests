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
		return englishHelloPrefix + "world"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == french {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("jobs", ""))
}
