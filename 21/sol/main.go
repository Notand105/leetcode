package main

import "fmt"

func main() {

	list1 := []int{1, 3, 5, 7, 9, 11, 12, 15, 19} //9
	list2 := []int{-9, -4, 2, 6, 9, 9, 9, 10}     // 8

	sol := mergeArr(list1, list2)
	fmt.Println(sol)
}

func mergeArr(list1, list2 []int) []int {
	//TODO comparar primer elemento de ambas listas, y eliminar ese de la lista inicial
	//Una vez la len de una lista sea igual a 0, añadir el resto de las que quedan en la otra lista
	// O investigar como implementar una estructura de datos similar a las listas enlazadas en golang
	//ver si existe alguna forma de OOP ver senio go en youtube para ver que ral

	aux := []int{}

	if len(list1) == 0 {
		return list2
	}
	if len(list2) == 0 {
		return list1
	}
	fmt.Println(list1, list2)
	for len(list1) > 0 {
		if len(list2) == 0 {
			list2 = append(list2, 101)
		} //los  numeros pueden estar entre -100 y 100, por eso si esta vacia l2 se continua el ciclo sin interferir
		if list1[0] < list2[0] {
			aux = append(aux, list1[0])
			list1 = list1[1:]
		} else {
			aux = append(aux, list2[0])
			list2 = list2[1:]
		}
	}

	return aux
}
