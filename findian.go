package main

import (
	"fmt"
	"strings"
	s "strings"
)

func main() {
	var inputString string
	fmt.Scan(&inputString)
	newString := strings.ToLower(inputString)
	if s.HasPrefix(newString, "i") {
		if s.Contains(newString, "a") {
			if s.HasSuffix(newString, "n") {
				print("Found!")
			} else {
				print("Not Found!")
			}
		} else {
			print("Not Found!")
		}
	} else {
		print("Not Found!")
	}
}
