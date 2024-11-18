package main

import "fmt"

const englishHelloPrefix = "hello "

func Hello(name, language string) string {
	if name == "" {
		return englishHelloPrefix + "world"
	}

	if language == "Spanish" {
		return "hola " + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("jobs", ""))
}
