package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0)
	var input int

	for i := 0; i < 10; i++ {
		fmt.Print("Enter an integer: ")
		fmt.Scanf("%d", &input)
		s = append(s, input)
	}

	BubbleSort(&s)
	fmt.Println(s)
}

//BubbleSort function
func BubbleSort(s *[]int) {
	for i := 0; i < 9; i++{
		for j :=0; j < 10-i-1; j++{
			// fmt.Println((*s)[i],(*s)[i+1])
			if (*s)[j] > (*s)[j+1]{
				Swap(s, j)
			}
		}	
	}
}

//Swap function
func Swap(s *[]int, pos int) {
	temp := (*s)[pos]
	(*s)[pos] = (*s)[pos+1]
	(*s)[pos+1] = temp
}
