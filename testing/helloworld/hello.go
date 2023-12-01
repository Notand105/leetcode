package main

import "fmt"

func main() {
	fmt.Print("here")
}

func HelloWorld(str, lang string) string {
	if len(str) == 0 {
		str = "world"
	}
	if lang == "es" {
		return "Hola " + str
	} else if lang == "it" {
		return "Ciao " + str
	}
	return "Hello " + str
}
