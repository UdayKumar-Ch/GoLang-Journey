package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var integersRequired string
	slic := make([]int, 0, 3)
	fmt.Println("Enter an integer")
	fmt.Scan(&integersRequired)
	for integersRequired != "X" {
		i, err := strconv.Atoi(integersRequired)
		slic = append(slic, i)
		if err == nil {
		}
		sort.Ints(slic)
		fmt.Println(slic)
		fmt.Scan(&integersRequired)
	}
}
