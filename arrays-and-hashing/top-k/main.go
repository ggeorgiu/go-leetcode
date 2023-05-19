package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{1}, 1))
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	c := make(map[int][]int)
	for num, freq := range m {
		c[freq] = append(c[freq], num)
	}

	var res []int
	for i := len(nums); i > 0; i-- {
		res = append(res, c[i]...)

		if len(res) == k {
			break
		}
	}

	return res
}
