package main

import (
	"fmt"
)

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (0.5*a*t*t + vo*t + so)
	}
	return fn
}

func main() {
	var a, vo, so, s float64
	fmt.Println("Enter value for acceleration: ")
	fmt.Scan(&a)
	fmt.Println("Enter value for velcoity: ")
	fmt.Scan(&vo)
	fmt.Println("Enter value for initial displacement: ")
	fmt.Scan(&so)
	fmt.Println("Values  : a", a, " vo:", vo, "so:", so)
	fn := GenDisplaceFn(a, vo, so)
	fmt.Println("Enter time to displacement : ")
	fmt.Scan(&s)
	fmt.Println("Displacement afer ", s, " seconds : ", fn(s))
}
