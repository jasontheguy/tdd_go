package main

import "fmt"

//Constant declarations for languages and prefix
const helloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

//Hello takes a name and a language string then concatenates
//a language specific greeting
func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	prefix := helloPrefix

	switch language {

	//Case can include type or not. This is sloppy tbh
	case "French":
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix

	}
	return prefix + name
}

func main() {
	fmt.Println("hello world")
}
