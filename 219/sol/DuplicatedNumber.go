package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	target := 2
	fmt.Println(duplicatedNum(nums, target))

}

func duplicatedNum(nums []int, target int) bool {

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == nums[i] && abs(i, j) <= target {
				return true
			}
		}
	}
	return false
}

func abs(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}
