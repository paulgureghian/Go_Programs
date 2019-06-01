// Created by Paul A. Gureghian in June 2019
// A Go source file which demonstrates the "zero value" in Go
package main

import "fmt"

var y string
var z int

func main() {

	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)

	y = "Bond, James Bond"

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T", z)
}
