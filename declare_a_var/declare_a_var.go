// Created by Paul A. Gureghian in May 2019
// A Go source file which explores declaring a var and reassigning it
package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)

	x = 99
	fmt.Println(x)

	y := 100 + 10
	fmt.Println(y)

	z := "Paul"
	fmt.Println(z)
}
