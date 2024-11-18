package main

import "fmt"

const englishHelloPrefix = "hello "

func Hello(name string) string {
	if name == "" {
		return englishHelloPrefix + "world"
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("jobs"))
}
