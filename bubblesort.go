package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	s "strings"
)

func BubbleSort(x []int) {
	end := len(x) - 1
	// fmt.Print("Array : ", *x)
	// fmt.Print("End : ", end)
	for {
		if end == 0 {
			break
		}
		for i := 0; i < len(x)-1; i++ {
			if (x)[i] > (x)[i+1] {
				(x)[i], (x)[i+1] = (x)[i+1], (x)[i]
			}
		}
		end -= 1
	}
}
func main() {
	inputReader := bufio.NewReader(os.Stdin)
	inputString, _ := inputReader.ReadString('\n')
	array := make([]int, 0)
	splitString := s.Fields(inputString)
	for _, val := range splitString {
		i, _ := strconv.Atoi(val)
		array = append(array, i)
	}

	BubbleSort(array)
	fmt.Print(array)
}
