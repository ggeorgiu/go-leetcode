package main

import "fmt"

func main() {
	fmt.Println(isValid("(]"))
}

func isValid(s string) bool {
	var stack []rune
	opposite := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, r := range s {
		headIdx := len(stack) - 1

		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		default:
			if len(stack) == 0 || stack[headIdx] != opposite[r] {
				return false
			}

			stack = stack[:headIdx]
		}
	}

	return len(stack) == 0
}
