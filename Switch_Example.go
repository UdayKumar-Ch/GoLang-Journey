package main

import "fmt"

func main() {
	var x int
	x = 0
	fmt.Printf("Running Switch Example\n")
	switch {
	case x > 1:
		fmt.Printf(" x value: %d", x)
		// break
	case x < 1:
		fmt.Printf("x value: %d", x)
		// break
	default:
		fmt.Printf("X value: %d", x)
		// break
	}
}
