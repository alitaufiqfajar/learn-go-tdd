package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const indonesian = "Indonesian"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const indonesianHelloPrefix = "Halo, "

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	prefix := ""

	switch language {
	case indonesian:
		prefix = indonesianHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}

func main() {
	response := Hello("Ali", "")

	fmt.Println(response)
}
