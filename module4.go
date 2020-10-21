package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (c Cow) Eat() {
	fmt.Println("Food Eaten :: Grass")
}

func (c Cow) Move() {
	fmt.Println("Locomotion Method :: Walk")
}

func (c Cow) Speak() {
	fmt.Println("Spoken Sound :: Moo")
}

func (b Bird) Eat() {
	fmt.Println("Food Eaten :: Worms")
}

func (b Bird) Move() {
	fmt.Println("Locomotion Method :: Fly")
}

func (b Bird) Speak() {
	fmt.Println("Spoken Sound :: Peep")
}

func (s Snake) Eat() {
	fmt.Println("Food Eaten :: Mice")
}

func (s Snake) Move() {
	fmt.Println("Locomotion Method :: Slither")
}

func (s Snake) Speak() {
	fmt.Println("Spoken Sound :: Hsss")
}

func main() {
	cow := Cow{}
	bird := Bird{}
	snake := Snake{}

	animals_created := make(map[string]Animal)

	fmt.Println("Choose one of the following options: newanimal (or) query")

	fmt.Println("Usage Instructions:")
	fmt.Println("newanimal <name_of_animal> <cow/bird/snake>")
	fmt.Println("query <name_of_animal> <eat/move/speak")

	fmt.Println("Example ::::")
	fmt.Println("newanimal lucky bird")
	fmt.Println("query lucky eat")
	inputReader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Print(">")
		inputString, _ := inputReader.ReadString('\n')
		splitString := s.Fields(inputString)

		if splitString[0] == "newanimal" {
			name_of_animal := splitString[1]
			switch {
			case "cow" == splitString[2]:
				animals_created[name_of_animal] = cow
			case "bird" == splitString[2]:
				animals_created[name_of_animal] = bird
			case "snake" == splitString[2]:
				animals_created[name_of_animal] = snake
			default:
				fmt.Println("New Animal can't be created. you need to choose one of the following: cow/snake/bird")
			}
		} else if splitString[0] == "query" {
			name_of_animal := splitString[1]

			if _, ok := animals_created[name_of_animal]; ok {
				fmt.Println("Animal presented")

			} else {
				continue
			}

			switch {
			case "eat" == splitString[2]:
				animals_created[name_of_animal].Eat()
			case "move" == splitString[2]:
				animals_created[name_of_animal].Move()
			case "speak" == splitString[2]:
				animals_created[name_of_animal].Speak()
			default:
				fmt.Println("Invalid option specified choose either eat/move/speak")
			}

		} else {
			fmt.Println("Used wrong syntax. Please follow Usage instructions.")
		}

	}
}
