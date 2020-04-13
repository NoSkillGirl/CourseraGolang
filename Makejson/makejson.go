package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	m := make(map[string]string)

	fmt.Print("Enter name: ")
	in1 := bufio.NewReader(os.Stdin)
	inputName, _ := in1.ReadString('\n')
	inputName = strings.TrimSuffix(inputName, "\n")

	fmt.Println("Enter address: ")
	in2 := bufio.NewReader(os.Stdin)
	inputAdd, _ := in2.ReadString('\n')
	inputAdd = strings.TrimSuffix(inputAdd, "\n")

	m["name"] = inputName
	m["address"] = inputAdd
	jsonString, _ := json.Marshal(m)
	fmt.Println(string(jsonString))
}
