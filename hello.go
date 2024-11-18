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

	prefix := getGreetingPrefix(language)

	return prefix + name
}

func getGreetingPrefix(language string) string {
	switch language {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func main() {
	fmt.Println(Hello("jobs", ""))
}
