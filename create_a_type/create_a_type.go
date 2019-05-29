// Created by Paul A. Gureghian in May 2019
// A Go source file to create data types
package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// Convert var b to type int
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
