package main

import "fmt"

func main() {
	// Type Conversions
	// Converting a variable from int32 to int16
	var x int32 = 1
	var y int16 = 2
	fmt.Printf("x value: %d\n", x)
	fmt.Printf("y value: %d\n", y)
	// If you run below code by uncommenting.
	// you will see an error.
	// x = y
	// This is how you need to typecast
	x = int32(y)
	fmt.Printf("x value after Type Conversion : %d", x)

	// Float32 - 6 digits of precision
	// Float64 - 15 digits of precision

	var a float64 = 1.2345123423422345
	var b float32 = 3.456
	fmt.Printf("Value of a: ", a, "\n")
	fmt.Printf("Value of b: ", b, "\n")
	b = float32(a)
	fmt.Printf("Value of b after type conversion: ", b, "\n")

	// Floating point/ Complex numbers
	var z complex128 = complex(2, 3)
	fmt.Printf("z value : ", z)

}
