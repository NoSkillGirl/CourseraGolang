package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	s := make([]int, 0, 3)
	for {
		fmt.Print("Enter a number: ")
		var input string
		fmt.Scanf("%s", &input)
		if input == "X" {
			break
		} else {
			value, err := strconv.Atoi(input)
			if err == nil {
				s = append(s, value)
				sort.Sort(sort.IntSlice(s))
				fmt.Println(s)
			}
		}
	}
}
