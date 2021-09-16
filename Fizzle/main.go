// Created by Paul A. Gureghian in Sept 2021.

// A game of Fizzle written in Go.

package main

import (
	"fmt"
	"math"
)

// Define function
func maxPoints(deck []int, k int) int {

	left := 0
	right := len(deck) - k
	var total, best int
	total = 0

	// k cards on the right give maxPoints.
	for i := right; i < len(deck); i++ {
		total += deck[i]

	}

	best = total

	// For loop runs 'k' times and tests all combos.
	for i := 0; i < k; i++ {

		// Remove points from right side and add to left side.
		total += deck[left] - deck[right]

		// Compare total points with current best points, keep the max.
		best = int(math.Max(float64(best), float64(total)))
		left++
		right++

	}

	return best

}

func main() {

	deck := []int{5, 3, 4, 4, 2, 3, 2, 6, 3}
	k := 4
	fmt.Println(maxPoints(deck, k))

}
