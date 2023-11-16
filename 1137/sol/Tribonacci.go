package main

import "fmt"

func main() {
	var in int = 0
	fmt.Printf("Ingresa el numero a calcular su tribonacci: ")
	fmt.Scan(&in)
	res := Tribonacci(in)
	fmt.Println(res)
}

func Tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	return Tribonacci(n-1) + Tribonacci(n-2) + Tribonacci(n-3)
}
