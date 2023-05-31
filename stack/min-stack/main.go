package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	top  int
	min  int
	vals []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{min: math.MaxInt32}
}

func (ms *MinStack) Push(val int) {
	ms.vals = append(ms.vals, val)
	ms.top = val

	if val <= ms.min {
		ms.mins = append(ms.mins, val)
		ms.min = val
	}
}

func (ms *MinStack) Pop() {
	if ms.top == ms.min {
		ms.mins = ms.mins[:len(ms.mins)-1]
		if len(ms.mins) > 0 {
			ms.min = ms.mins[len(ms.mins)-1]
		} else {
			ms.min = math.MaxInt32
		}
	}

	ms.vals = ms.vals[:len(ms.vals)-1]
	if len(ms.vals) > 0 {
		ms.top = ms.vals[len(ms.vals)-1]
	} else {
		ms.top = 0
	}
}

func (ms *MinStack) Top() int {
	return ms.top
}

func (ms *MinStack) GetMin() int {
	return ms.min
}

func main() {
	minStack := Constructor()
	minStack.Push(2147483646)
	minStack.Push(2147483646)
	minStack.Push(2147483647)
	minStack.Top()
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	minStack.Push(2147483647)
	minStack.Top()
	fmt.Println(minStack.GetMin())
	minStack.Push(-2147483648)
	minStack.Top()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())

}
