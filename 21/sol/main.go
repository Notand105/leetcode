package main

import "fmt"

func main() {

	list1 := []int{1, 3, 5, 7, 9, 11, 12, 15, 19} //9
	list2 := []int{2, 4, 6, 9, 9, 9, 9, 10}       // 8

	sol := mergeArr(list1, list2)
	fmt.Println(sol)
}

func mergeArr(list1, list2 []int) []int {

	aux := []int{}
	j := 0
	for i := 0; i < len(list1) || j < len(list2); i++ {
		if i >= len(list1) && j < len(list2) {
			i--
		} else if j >= len(list2) && i < len(list1) {
			j--
		}

		fmt.Printf("{ i: %d , j: %d , list1[i]: %d, list2[j]: %d } \n", i, j, list1[i], list2[j])
		if list1[i] < list2[j] {
			aux = append(aux, list1[i])
			fmt.Printf("agregamos %d \n", list1[i])
		} else {
			aux = append(aux, list2[j])
			fmt.Printf("agregamos %d \n", list2[j])
			j++
			i--
		}
	}

	return aux
}
