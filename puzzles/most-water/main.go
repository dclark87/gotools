package main

import "fmt"

// Problem courtesy of:
// https://leetcode.com/problems/container-with-most-water/
func mostWater(nums []int) int {
	// Naive approach, get every possible size: O(n^2)
	// Faster? start on the outter-most and work inward.
	var maxVol, vol = 0, 0
	var left = 0
	var right = len(nums) - 1
	for right-left > 0 {
		if nums[left] > nums[right] {
			vol = (right - left) * nums[right]
			right--
		} else {
			vol = (right - left) * nums[left]
			left++
		}
		if vol > maxVol {
			maxVol = vol
		}
	}

	return maxVol
}

func main() {
	var tcs [][]int
	tcs = append(tcs, []int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	for _, tc := range tcs {
		fmt.Printf("input: %+v\n", tc)
		fmt.Println("output:", mostWater(tc))
	}
}
