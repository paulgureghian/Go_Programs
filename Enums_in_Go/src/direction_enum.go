package main

import "fmt"

// Direction - Custom type to hold value for week day ranging from 1-4
type Direction int

// Declare related constants for each direction starting with index 1
const (
	North Direction = iota + 1 // EnumIndex = 1
	East                       // EnumIndex = 2
	South                      // EnumIndex = 3
	West                       // EnumIndex = 4
)

// String - Creating common behavior - give the type a String function
func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex functio
func (d Direction) EnumIndex() int {
	return int(d)
}

func main() {
	var d Direction = West
	fmt.Println(d)             // Print : West
	fmt.Println(d.String())    // Print : West
	fmt.Println(d.EnumIndex()) // Print : 4
}
