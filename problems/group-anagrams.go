package problems

import (
	"fmt"
	"sort"
)

//medium

//https://leetcode.com/problems/group-anagrams/

func GroupAnagrams(strs []string) [][]string {

	m := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		flag := sortString(strs[i])
		m[flag] = append(m[flag], strs[i])
	}
	var result [][]string
	for _, val := range m {
		result = append(result, val)
	}
	fmt.Println(result)
	return result
}

func sortString(s string) string {
	runes := []rune(s) // handle Unicode properly
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
