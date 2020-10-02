package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	s "strings"
)

func Swap(x []int, i int) {
	(x)[i], (x)[i+1] = (x)[i+1], (x)[i]
}

func BubbleSort(x []int) {
	end := len(x) - 1
	for j := len(x) - 1; j >= 0; j-- {
		for i := 0; i < j; i++ {
			if (x)[i] > (x)[i+1] {
				Swap(x, i)
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
