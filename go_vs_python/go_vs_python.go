// Created by Paul A. Gureghian in May 2019
// A Go source file which has some facts about Go and Python
package main

import "fmt"

func main() {
	fmt.Println("Hi, I'm Paul")

	foo()

	fmt.Println("not loosely typed")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("Go is a strongly typed language")
}

func bar() {
	fmt.Println("like Python")
}
