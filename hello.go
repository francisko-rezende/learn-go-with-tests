package main

import "fmt"

const (
	englishHelloPrefix = "hello "
	spanishHelloPrefix = "hola "
	spanish            = "Spanish"
)

func Hello(name, language string) string {
	if name == "" {
		return englishHelloPrefix + "world"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("jobs", ""))
}
