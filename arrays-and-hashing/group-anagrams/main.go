package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string, len(strs))

	for _, v := range strs {
		key := sortLetters(v)
		if val, ok := m[key]; ok {
			m[key] = append(val, v)
			continue
		}

		m[key] = []string{v}
	}

	res := make([][]string, len(m))
	i := 0
	for _, vals := range m {
		res[i] = vals
		i++
	}

	return res
}

func sortLetters(s string) string {
	let := strings.Split(s, "")
	sort.Strings(let)
	return strings.Join(let, "")
}
