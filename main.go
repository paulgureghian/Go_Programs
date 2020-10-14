// Created by Paul A. Gureghian in Oct 2020
// This Go program demonstrates importing different versions of the same package

package main

import (
	"fmt"

	"rsc.io/quote"
	quotev2 "rsc.io/quote/v2"
)

func main() {
	
	fmt.Println(quote.Hello())
	fmt.Println(quotev2.HelloV2())
}
