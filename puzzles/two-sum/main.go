package main

import "fmt"

// Problem courtesy of:
// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	// Build map of {num : [idx, ...]}.
	m := make(map[int]int)
	var ans []int
	for i, num := range nums {
		diff := target - num
		// If difference is found in map, return key and current idx.
		if val, ok := m[diff]; ok {
			ans = append(ans, val)
			ans = append(ans, i)
			break
		} else {
			m[num] = i
		}
	}
	return ans
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
