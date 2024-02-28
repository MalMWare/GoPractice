package main

import "fmt"

// Basics
// func main() {
// 	fmt.Println("Hello, world")
// }

//How to test the code
const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

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

	// lets change this swithc to be more inclusive
	// prefix := englishHelloPrefix
	// switch language {
	// case "French":
	// 	prefix = frenchHelloPrefix
	// case "Spanish":
	// 	prefix = spanishHelloPrefix
	// }
	// return prefix + name
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
