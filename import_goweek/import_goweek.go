// Created by Paul A. Gureghian in May 2019
// A Go source file which imports and uses the package goweek
package main

import "fmt"
import "github.com/paulgureghian/Go/goweek"

func main() {
	week, goweek_err := goweek.NewWeek(2015, 33)

	n, print_err := fmt.Println(week)
	n1, print_err1 := fmt.Println(goweek_err)

	fmt.Println(n)
	fmt.Println(print_err)

	fmt.Println(n1)
	fmt.Println(print_err1)
}
