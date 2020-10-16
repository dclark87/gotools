package main

import (
	"fmt"
	"math"
	"math/rand"
)

/**
You have n fair coins and you flip them all at the same time.
Any that come up tails you set aside. The ones that come up heads you flip
again. How many rounds do you expect to play before only one coin remains?

Write a function that, given n, returns the number of rounds you'd expect to
play until one coin remains.
**/

// Method 1 - do this using probability.
func numRounds(n int) int {
	rand.Seed(42)
	var r, head int
	for n > 1 {
		var nn = 0
		for i := 0; i < n; i++ {
			head = rand.Intn(2)
			// tails, remove from next round
			if head == 0 {
				nn++
			}
		}
		r++
		n = n - nn
	}
	return r
}

// Method 2 - do this using recursive divide by 2 (log base 2).
func numRounds2(n int) int {
	return int(math.Log2(float64(n)))
}

// Given an array of integers, determine whether it contains a Pythagorean triplet.
// Recall that a Pythogorean triplet (a, b, c) is defined by the equation a^2+ b^2= c^2.
func main() {
	fmt.Println("vim-go")
}
