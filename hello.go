package main

import "fmt"

// Basics
// func main() {
// 	fmt.Println("Hello, world")
// }

//How to test the code

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
