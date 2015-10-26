package main

import (
	"fmt"
	"os"

	"github.com/icmpecho/hello"
)

func main() {
	name := "icmpecho"
	if len(os.Args) >= 2 {
		name = os.Args[1]
	}
	greet := hello.Greet(name)
	fmt.Printf(greet)
}
