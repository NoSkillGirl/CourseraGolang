package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Print("Enter a number: ")
	var input string
	fmt.Scanf("%s", &input)
	i, err := strconv.ParseFloat(input, 64)
	if err == nil {
		var y int = int(i)
		fmt.Println(y)
	} else {
		fmt.Println("Number entered is not valid")
	}
}
