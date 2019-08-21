package main

import "fmt"

const french = "French"
const italian = "Italian"
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const italianHelloPrefix = "Ciao, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case italian:
		prefix = italianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Franco", ""))
}
