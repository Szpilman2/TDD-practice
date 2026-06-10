package main

import "fmt"

const englishHelloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"
const frenchHelloPrefix = "Bonjour,"

const spanish = "spanish"
const french = "french"

func Hello(name string, language string) string {
	if name == "" {
		return "Hello, World"
	}

	prefix := greetingPrefix(language)

	return fmt.Sprintf("%s %s", prefix, name)
}

func greetingPrefix(language string) string{
	prefix := englishHelloPrefix
	switch language {
		case spanish:
			prefix = spanishHelloPrefix
		case french:
			prefix = frenchHelloPrefix
	}

	return prefix
}

func main() {
	fmt.Println(Hello("World", ""))
}