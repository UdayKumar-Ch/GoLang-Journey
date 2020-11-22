package main

import "fmt"

// Module 2 Activity trunc.go
func main() {
	var input_number float64
	// Reading a float64 number
	fmt.Printf("Enter Number: ")
	fmt.Scan(&input_number)
	fmt.Print("Result After conversion to int: ", int64(input_number))
}
