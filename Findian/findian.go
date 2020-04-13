package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a string: ")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	fmt.Println(input)
	if err == nil {
		fc := string(input[0])
		l := len(string(input)) - 2
		lc := string(input[l])
		if (fc == "i" || fc == "I") && (lc == "n" || lc == "N") && strings.ContainsAny(input, "a | A") {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	} else {
		fmt.Println("Error occured")
	}
}
