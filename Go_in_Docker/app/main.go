package main

import (
	"fmt"
	"github.com/paulgureghian/Go_Programs/Go_in_Docker/package/test"
)

func main() {
	message := "Call Package:"
	fmt.Println(message, test.Test())
}
