package main

import (
	"fmt"
	"math"
)

// Problem courtesy of:
// https://leetcode.com/problems/reverse-integer/

// revint reverses any 32-bit signed integer.
// For example, 1234 -> 4321.
// Must ensure non-overflow:
// [-2^31, 2^31-1]
func revint(num int) int {
	var rev = 0
	var pop int
	for num != 0 {
		pop = num % 10 // 4
		num /= 10      // 123.4
		// Large positive number.
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7) {
			return 0
			// Large negative number.
		} else if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}

	return rev
}

func main() {
	var tcs []int
	tcs = append(tcs, 12345)
	tcs = append(tcs, 7463847412)
	tcs = append(tcs, -1)
	tcs = append(tcs, 8463847412)
	tcs = append(tcs, -8463847412)
	tcs = append(tcs, -9463847412)
	fmt.Println("maxint:", math.MaxInt32)
	for _, d := range tcs {
		fmt.Println("input: ", d)
		fmt.Println("output: ", revint(d))
	}
}
