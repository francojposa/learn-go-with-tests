package main

import "fmt"

const italian = "Italian"
const englishHelloPrefix = "Hello, "
const italianHelloPrefix = "Ciao, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == italian {
		return italianHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Franco", ""))
}
