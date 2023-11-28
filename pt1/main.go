package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(findLargest([]int{3, 5, 4, 6, 9, 8, 7}))
	fmt.Println(findLargest([]int{-3, -1, 5, 4, 6, -9, 8, 7}))
	fmt.Println(findLargest([]int{}))
}

func findLargest(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("slice is empty")
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max, nil
}
