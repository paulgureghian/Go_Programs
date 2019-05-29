// Created by Paul A. Gureghian in May 2019
// A Go souce file which explores the 'Println' function
package main

import "fmt"

// The Println() can take a mixture of different data types at the same time
// The 'n' holds the returned value of bytes, the _ throws away the returned error, if any

func main() {
	n, _ := fmt.Println("Hello, playground", 42, true)
	fmt.Println(n)
}
