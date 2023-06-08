package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 7}))
}

func containsDuplicate(nums []int) bool {
	c := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := c[nums[i]]; ok {
			return true
		}
		c[nums[i]]++
	}

	return false
}
