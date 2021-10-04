package main

import (
	"fmt"
	"time"
)

func numSearch(targetVal int) {

	searchlist := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	time.Sleep((time.Duration(targetVal) * 100) * time.Millisecond)

	for _, num := range searchlist {

		if num == targetVal {

			fmt.Printf("Found the targetVal: %d \n", targetVal)
			return
		}
	}

	fmt.Println("TargetVal not found : (")

}

func main() {

	go numSearch(10)
	go numSearch(1)

	time.Sleep(10000 * time.Millisecond)
	fmt.Println("Program has finished")

}
