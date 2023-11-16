package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 6}
	res := missingNumber(arr)
	fmt.Println("El numero que falta en el array es: ", res)
}

func missingNumber(arr []int) int {
	//we need to return the only number that´s not there,
	//so we sum the number from 1 to len(arr) ant then substract the sum of the array
	var arrlen int = len(arr)
	var total int = (arrlen * (arrlen + 1)) / 2
	var arrsum int = 0
	for i := 0; i < len(arr); i++ {
		arrsum += arr[i]
	}
	fmt.Println(arrlen, total, arrsum, total-arrsum)
	return total - arrsum
}
