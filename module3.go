package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

type Animal struct {
	food_eat          string
	locomotion_method string
	spoken_sound      string
}

func (a Animal) eat() string {
	return a.food_eat
}

func (a Animal) move() string {
	return a.locomotion_method
}

func (a Animal) speak() string {
	return a.spoken_sound
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}
	// eat move speak
	inputReader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print(">")
		inputString, _ := inputReader.ReadString('\n')
		splitString := s.Fields(inputString)
		if splitString[0] == "cow" {
			if splitString[1] == "eat" {
				fmt.Println(cow.eat())
			} else if splitString[1] == "move" {
				fmt.Println(cow.move())
			} else if splitString[1] == "speak" {
				fmt.Println(cow.speak())
			} else {
				fmt.Println("second string is not valid : you need to choose either eat or move or speak")
			}
		} else if splitString[0] == "bird" {
			if splitString[1] == "eat" {
				fmt.Println(bird.eat())
			} else if splitString[1] == "move" {
				fmt.Println(bird.move())
			} else if splitString[1] == "speak" {
				fmt.Println(bird.speak())
			} else {
				fmt.Println("second string is not valid : you need to choose either eat or move or speak")
			}
		} else if splitString[0] == "snake" {
			if splitString[1] == "eat" {
				fmt.Println(snake.eat())
			} else if splitString[1] == "move" {
				fmt.Println(snake.move())
			} else if splitString[1] == "speak" {
				fmt.Println(snake.speak())
			} else {
				fmt.Println("second string is not valid : you need to choose either eat or move or speak")
			}
		} else {
			fmt.Println("First string is invalid : you need to choose either cow or bird or snake")
		}
	}
}
