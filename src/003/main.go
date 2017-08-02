package main

import (
	"fmt"
)

func main() {
	j, k := 0, 0
	fmt.Println(j + k)
	lengthOfLongestSubstring("pwkpww")
}

func lengthOfLongestSubstring(s string) int {
	var maxLen, i int
	runes := []rune(s)
	index := [128]int{}
	for j := 0; j < len(runes); j++ {
		i = max(index[runes[j]], i)
		maxLen = max(maxLen, j-i+1)
		index[runes[j]] = j + 1
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
