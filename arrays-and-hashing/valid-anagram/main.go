package main

import "fmt"

func main() {
	fmt.Println(isAnagram("anagram", "naagram"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	c1 := make(map[uint8]int, len(s))
	c2 := make(map[uint8]int, len(s))
	for i := 0; i < len(s); i++ {
		c1[s[i]]++
		c2[t[i]]++
	}

	for k, v := range c1 {
		if c2[k] != v {
			return false
		}
	}

	return true
}
