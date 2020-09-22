package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	var fileName string
	fmt.Print("Enter filename : ")
	fmt.Scan(&fileName)

	textLines := make([]Person, 0)

	data, _ := os.Open(fileName)
	defer data.Close()

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()

		result := s.Split(line, " ")
		fName := result[0]
		lName := result[1]
		if len(result[0]) > 20 {
			fName = fName[:20]
		}
		if len(result[1]) > 20 {
			lName = lName[:20]
		}

		p := Person{fname: fName, lname: lName}

		textLines = append(textLines, p)
	}

	fmt.Println("Following are the Lines:")

	for _, person := range textLines {
		fmt.Println(" fName : ", person.fname, " lName :", person.lname)
	}

}
