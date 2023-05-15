package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}

func containsDuplicate(nums []int) bool {
	c := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if c[nums[i]] {
			return true
		}
		c[nums[i]] = true
	}

	return false
}
