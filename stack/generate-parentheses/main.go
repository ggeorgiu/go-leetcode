package main

import (
	"fmt"
	"strings"
)

var stack []string
var res []string

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	stack = []string{}
	res = []string{}

	back(0, 0, n)
	return res
}

func back(open, closed, n int) {
	if open == n && closed == n {
		res = append(res, strings.Join(stack, ""))
		return
	}

	if open < n {
		stack = append(stack, "(")
		back(open+1, closed, n)
		stack = stack[:len(stack)-1]
	}
	if closed < open {
		stack = append(stack, ")")
		back(open, closed+1, n)
		stack = stack[:len(stack)-1]
	}
}
