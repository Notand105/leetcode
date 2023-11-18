package main

import (
	"fmt"
	"strings"
)

func main() {

	var str string = ""
	cleanStr(&str)
	fmt.Println("cadena a analizar: ", str)
	var res bool = validatePalindrome(str)

	fmt.Println(res)

}

func cleanStr(str *string) {
	*str = strings.ToLower(*str)
	var aux string = *str
	var aux2 string

	for i := 0; i < len(aux); i++ {
		if isInArr(aux[i]) {
			aux2 = aux2 + string(aux[i])
		}
	}
	*str = aux2
}

func isInArr(s byte) bool {
	allow := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	for i := 0; i < len(allow); i++ {
		if allow[i] == s {
			return true
		}
	}
	return false
}

func validatePalindrome(str string) bool {

	var aux int = len(str) - 1

	for i := 0; i < len(str) && i <= aux; i++ {
		if str[i] != str[aux] {
			return false
		}
		aux--
	}

	return true
}
