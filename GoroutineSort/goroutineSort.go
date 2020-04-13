package main

import (
	"fmt"
	"sort"
	"strconv"
)

//function to sort a given slice
func sortSlice(s []int, c chan []int) {
	sort.Sort(sort.IntSlice(s))
	fmt.Println("List sorted by goroutine: ", s)
	c <- s
}

func main() {
	s := make([]int, 0, 3)
	//saving the input in a slice
	for {
		fmt.Print("Enter a number or to exit enter X: ")
		var input string
		fmt.Scanf("%s", &input)
		if input == "X" || input == "x" {
			break
		} else {
			value, err := strconv.Atoi(input)
			if err == nil {
				s = append(s, value)
			}
		}
	}

	//deviding the slice in 4 equal parts
	divLen := len(s) / 4
	rem := len(s) % 4
	var s1, s2, s3, s4 []int
	var lens1, lens2, lens3 int

	//finding the lengths of each sub slice according to reminder
	if rem == 0 {
		lens1 = divLen
		lens2 = divLen * 2
		lens3 = divLen * 3
	} else if rem == 1 {
		lens1 = divLen + 1
		lens2 = divLen*2 + 1
		lens3 = divLen*3 + 1
	} else if rem == 2 {
		lens1 = divLen + 1
		lens2 = divLen*2 + 2
		lens3 = divLen*3 + 2
	} else if rem == 3 {
		lens1 = divLen + 1
		lens2 = divLen*2 + 2
		lens3 = divLen*3 + 3
	}

	//creating 4 sub slice according to the length obtained from above
	s1 = s[:lens1]
	s2 = s[lens1:lens2]
	s3 = s[lens2:lens3]
	s4 = s[lens3:]

	//goroutine functions
	c := make(chan []int)
	go sortSlice(s1, c)
	go sortSlice(s2, c)
	go sortSlice(s3, c)
	go sortSlice(s4, c)

	//geting results from chanel
	sortedS1 := <-c
	sortedS2 := <-c
	sortedS3 := <-c
	sortedS4 := <-c

	//merging the slice sorted sub slice into one slice
	mergeSlice := make([]int, len(s))
	for i, v := range sortedS1 {
		mergeSlice[i] = int(v)
	}
	for i, v := range sortedS2 {
		mergeSlice[i+len(sortedS1)] = int(v)
	}
	for i, v := range sortedS3 {
		mergeSlice[i+len(sortedS1)+len(sortedS2)] = int(v)
	}
	for i, v := range sortedS4 {
		mergeSlice[i+len(sortedS1)+len(sortedS2)+len(sortedS3)] = int(v)
	}

	//sorting and printing the merged slice
	sort.Sort(sort.IntSlice(mergeSlice))
	fmt.Println("Final sorted list: ", mergeSlice)
}
