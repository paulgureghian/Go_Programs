// Created by Paul A. Gureghian in May 2019
// A Go source file which demonstrates control flow
// With the 3 methods of: Sequence, Loops, and Conditions
package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, I'm Paul Gureghian.")

	foo()

	fmt.Println("and I like it")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm learning the Go language")
}

func bar() {
	fmt.Println("and I'm good at it.")
}
