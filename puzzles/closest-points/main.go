package main

import (
	"fmt"
	"math"
)

/**
Given a set of points (x, y) on a 2D cartesian plane, find the two closest
points. For example, given the points
[(1, 1), (-1, -1), (3, 4), (6, 1), (-1, -6), (-4, -3)],
return [(-1, -1), (1, 1)].
**/

type coord struct {
	x, y int
}

func distance(p1, p2 coord) float64 {
	var a = float64(p1.x - p2.x)
	var b = float64(p1.y - p2.y)
	return math.Sqrt(a*a + b*b)
}

func closestPoints(coords []coord) []coord {
	var i = 0
	var min = math.MaxFloat64
	var closest = make([]coord, 2)
	for i < len(coords) {
		for j := i + 1; j < len(coords); j++ {
			if m := distance(coords[i], coords[j]); m < min {
				min = m
				closest[0] = coords[i]
				closest[1] = coords[j]
			}
		}
		i++
	}
	return closest
}

func main() {
	fmt.Println("vim-go")
}
