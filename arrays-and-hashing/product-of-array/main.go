package main

import "fmt"

func main() {
    fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
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
