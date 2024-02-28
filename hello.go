package main

import "fmt"

// Basics
// func main() {
// 	fmt.Println("Hello, world")
// }

//How to test the code

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
