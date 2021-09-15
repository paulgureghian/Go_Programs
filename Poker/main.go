// Created by Paul A. Gureghian in Sept 2021.

// A game of Poker written in Go.

package main

import (
	"fmt"
	"sort"
)

// Define function.
func isHandOfStraights(hand []int, k int) bool {

	// Check if hand is divisable by k.
	if len(hand)%k != 0 {
		return false
	}

	// Count the occurance of each card in the hand.
	count := make(map[int]int)
	for _, i := range hand {
		count[i] = count[i] + 1
	}

	// Sort the list.
	sort.Ints(hand)
	i := 0
	n := len(hand)

	// Nested loop that runs k times.
	for i < n {
		current := hand[i]
		for j := 0; j < k; j++ {
			
			// Check if the current card and the next k-1 cards are in the count map.
			if _, ok := count[current+j]; !ok || count[current+j] == 0 {
				return false
			}

			// When a card is in the count map, decrement it.
			count[current+j]--
		}

		// Find the next group's smallest card.
		for i < n && count[hand[i]] == 0 {
			i++
		}
	}

	// If all cards are sorted into groups.
	return true
}

func main() {

	hand := []int{5, 2, 4, 4, 1, 3, 5, 6, 3}
	k := 3
	fmt.Println(isHandOfStraights(hand, k))

	hand2 := []int{1, 9, 3, 5, 7, 4, 2, 9, 11}
	k = 2
	fmt.Println(isHandOfStraights(hand2, k))

}
