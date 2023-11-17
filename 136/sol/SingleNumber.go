package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 1, 5, 6, 2, 3, 4, 6}
	fmt.Println(SingleNumber(nums))
}

func SingleNumber(nums []int) int {
	var aux int

	for i := 0; i < len(nums); i++ {
		aux ^= nums[i]
	}
	return aux

}
