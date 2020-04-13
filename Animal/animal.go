package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Animal struct
type Animal struct {
	food       string
	locomotion string
	noise      string
}

//Eat func
func (a *Animal) Eat() {
	fmt.Println(a.food)
}

//Move func
func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

//Speak func
func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	cow := Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}

	bird := Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}

	snake := Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	fmt.Println("Enter animal name and information request about the animal (in case you want to exit type 'X'): ")
	for {
		fmt.Print(">")
		in1 := bufio.NewReader(os.Stdin)
		input, _ := in1.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		//exit if user types x or X
		if input == "X" || input == "x" {
			break
		}

		s := strings.Split(input, " ")
		name, info := s[0], s[1]

		if name == "cow" {
			if info == "eat" {
				cow.Eat()
			}
			if info == "move" {
				cow.Move()
			}
			if info == "speak" {
				cow.Speak()
			}
		}

		if name == "bird" {
			if info == "eat" {
				bird.Eat()
			}
			if info == "move" {
				bird.Move()
			}
			if info == "speak" {
				bird.Speak()
			}
		}

		if name == "snake" {
			if info == "eat" {
				snake.Eat()
			}
			if info == "move" {
				snake.Move()
			}
			if info == "speak" {
				snake.Speak()
			}
		}
	}
}
