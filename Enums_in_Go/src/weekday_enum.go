package main

import "fmt"

// Weekday - Custom type to hold value for weekday ranging from 1-7
type Weekday int

// Declare related constants for each weekday starting with index 1
const (
	Sunday    Weekday = iota + 1 // EnumIndex = 1
	Monday                       // EnumIndex = 2
	Tuesday                      // EnumIndex = 3
	Wednesday                    // EnumIndex = 4
	Thursday                     // EnumIndex = 5
	Friday                       // EnumIndex = 6
	Saturday                     // EnumIndex = 7
)

// String - Creating common behavior - give the type a String function
func (w Weekday) String() string {
	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[w-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (w Weekday) EnumIndex() int {
	return int(w)
}

func main() {
	var weekday = Sunday
	fmt.Println(weekday)             // Print : Sunday
	fmt.Println(weekday.String())    // Print : Sunday
	fmt.Println(weekday.EnumIndex()) // Print : 1
}
