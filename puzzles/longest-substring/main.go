package main

import (
	"fmt"
)

// Problem courtesy of:
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	pos := make(map[rune]int)
	ans := 0
	beg := 0
	for end, r := range s {
		if p := pos[r]; p > beg {
			beg = p
		}
		// Store idx+1 to make up for +1 in `cur` calc.
		// This is needed in the case of a single char string.
		pos[r] = end + 1
		if cur := end - beg + 1; cur > ans {
			ans = cur
		}
	}
	return ans
}

func main() {
	in := "c"
	length := lengthOfLongestSubstring(in)
	fmt.Printf("%s: %d\n", in, length)
	in = "cdlfjcd"
	length = lengthOfLongestSubstring(in)
	fmt.Printf("%s: %d\n", in, length)
}
