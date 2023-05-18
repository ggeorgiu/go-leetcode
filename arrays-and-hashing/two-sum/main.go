package main

import "fmt"

func main() {
	fmt.Println(twoSum2([]int{1, 2, 3, 4, 5}, 9))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if v, ok := m[diff]; ok {
			return []int{i, v}
		}

		m[nums[i]] = i
	}

	return []int{}
}
