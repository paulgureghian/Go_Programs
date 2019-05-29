// Created by Paul A. Gureghian in May 2019
// A Go source file which explores the 'type' keyword
package main

import (
	"fmt"
)

var y = 42
var z string = "Shaken, not stirred"
var a string = `James said, "Shaken,not stirred"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
