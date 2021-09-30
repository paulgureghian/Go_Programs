// Created by Paul A. Gureghian in Sept 2021.

// A simple port scanner written in Go.

package main

import (
	"fmt"
	"os"

	"github.com/paulgureghian/Go_Programs/Port_Scanner/portscan"
)

func main() {

	fmt.Println("Checking for available ports...")
	ports := portscan.PortScan(os.Args [1])

	fmt.Println("Ports available: ", ports)

}
