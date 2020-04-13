package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Animal interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

//Cow type
type Cow struct {
}

//Eat func
func (c *Cow) Eat() {
	fmt.Println("grass")
}

//Move func
func (c *Cow) Move() {
	fmt.Println("walk")
}

//Speak func
func (c *Cow) Speak() {
	fmt.Println("moo")
}

//Bird type
type Bird struct {
}

//Eat func
func (b *Bird) Eat() {
	fmt.Println("worms")
}

//Move func
func (b *Bird) Move() {
	fmt.Println("fly")
}

//Speak func
func (b *Bird) Speak() {
	fmt.Println("peep")
}

//Snake type
type Snake struct {
}

//Eat func
func (s *Snake) Eat() {
	fmt.Println("mice")
}

//Move func
func (s *Snake) Move() {
	fmt.Println("slither")
}

//Speak func
func (s *Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animalMap := make(map[string]Animal)

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

		str := strings.Split(input, " ")
		cmd, name, typeInfo := str[0], str[1], str[2]

		if cmd == "newanimal" {
			// if typeInfo == "cow" {
			// 	animalMap[name] = &Cow{}
			// }
			switch typeInfo {
			case "cow":
				animalMap[name] = &Cow{}
				fmt.Println("Created it")
			case "bird":
				animalMap[name] = &Bird{}
				fmt.Println("Created it")
			case "snake":
				animalMap[name] = &Snake{}
				fmt.Println("Created it")
			default:
				fmt.Println("Entered type is not correct")
			}
		}
		if cmd == "query" {
			switch typeInfo {
			case "eat":
				animalMap[name].Eat()
			case "move":
				animalMap[name].Move()
			case "speak":
				animalMap[name].Speak()
			default:
				fmt.Println("Name or method is/are not correct")
			}
		}
	}
}
