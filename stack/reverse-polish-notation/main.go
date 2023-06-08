package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}

func evalRPN(tokens []string) int {
	var st []int

	for _, t := range tokens {
		switch t {
		case "+", "-", "/", "*":
			a := st[len(st)-2]
			b := st[len(st)-1]

			res := op(a, b, t)
			st = st[:len(st)-2]
			st = append(st, res)
		default:
			nr, _ := strconv.Atoi(t)
			st = append(st, nr)
		}
	}

	return st[0]
}

func op(a int, b int, t string) int {
	switch t {
	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
		return a / b
	case "*":
		return a * b
	}

	return -1
}
