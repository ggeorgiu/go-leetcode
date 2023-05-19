package main

import "fmt"

func main() {
	fmt.Println(productExceptSelfOnce([]int{2, 3, 5, 0}))
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	prefix := 1
	for i, n := range nums {
		res[i] = prefix
		prefix *= n
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}

func productExceptSelfOnce(nums []int) []int {
	l := len(nums)
	res := make([]int, l)

	prefix := 1
	postfix := 1
	for i := 0; i < l; i++ {
		j := l - i - 1
		if res[i] == 0 && postfix != 0 {
			res[i] = 1
		}
		if res[j] == 0 && prefix != 0 {
			res[j] = 1
		}

		res[i] *= prefix
		res[j] *= postfix

		prefix *= nums[i]
		postfix *= nums[j]
	}

	return res
}
