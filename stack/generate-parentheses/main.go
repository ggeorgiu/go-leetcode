package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var stack []string
	var res []string
	var back func(int, int)

	back = func(open int, closed int) {
		if open == n && closed == n {
			res = append(res, strings.Join(stack, ""))
			return
		}

		if open < n {
			stack = append(stack, "(")
			back(open+1, closed)
			stack = stack[:len(stack)-1]
		}
		if closed < open {
			stack = append(stack, ")")
			back(open, closed+1)
			stack = stack[:len(stack)-1]
		}
	}

	back(0, 0)

	return res
}
