package main

import (
	"fmt"
)

func main() {

	n := []int{11, 12, 0, 13, 14, 15, 16, 3, 7, 2, 5, 8, 4, 6, 1}
	fmt.Println(longestConsecutive(n))
}

type interval struct {
	start int
	end   int
}

func longestConsecutive(nums []int) int {
	m := make(map[int]interval, len(nums))
	var max int

	for _, v := range nums {
		if _, ok := m[v]; ok {
			continue
		}

		in := interval{v, v}
		lt, oklt := m[v-1]
		if oklt {
			in.start = lt.start
		}
		gt, okgt := m[v+1]
		if okgt {
			in.end = gt.end
		}

		length := in.end - in.start + 1
		if length > max {
			max = length
		}

		if oklt {
			m[v-1] = in
		}
		if okgt {
			m[v+1] = in
		}

		m[v] = in
		m[in.start] = in
		m[in.end] = in
	}

	return max
}
