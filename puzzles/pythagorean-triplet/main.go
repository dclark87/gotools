package main

import (
	"fmt"
	"math"
	"sort"
)

func containsTriplet(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	// sort the array (in-place): [3,4,5,6]
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	// Make a set (hash map).
	var m = make(map[int]bool)
	for _, a := range arr {
		m[a] = true
	}

	// Iterate from largest to smallest
	var i = len(arr) - 1
	var delta int
	for i > 0 {
		// If largest is 0, not valid.
		if arr[i] == 0 {
			return false
		}
		delta = arr[i]*arr[i] - arr[i-1]*arr[i-1]
		// if delta is not evenly square-root-able, check next element.A
		// else
		// check if sqrt(delta) is in map, if so then we found one. continue
		if sqrt := math.Sqrt(float64(delta)); sqrt == math.Floor(sqrt) {
			if _, ok := m[int(math.Sqrt(float64(delta)))]; ok {
				return true
			}
		}
		i--
	}
	return false
}

// Given an array of integers, determine whether it contains a Pythagorean triplet.
// Recall that a Pythogorean triplet (a, b, c) is defined by the equation a^2+ b^2= c^2.
func main() {
	fmt.Println("vim-go")
}
