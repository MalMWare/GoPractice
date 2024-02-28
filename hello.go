package main

import "fmt"

// Basics
// func main() {
// 	fmt.Println("Hello, world")
// }

//How to test the code
const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	//this could be a switch!
	// if language == spanish {
	// 	return spanishHelloPrefix + name
	// }

	// if language == french {
	// 	return frenchHelloPrefix + name
	// }
	// return englishHelloPrefix + name

	prefix := englishHelloPrefix

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
