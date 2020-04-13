package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//Person struct
type person struct {
	fname string
	lname string
}

func main() {
	fmt.Println("Enter the file path: ")
	var filePath string
	fmt.Scanf("%s", &filePath)

	dat, e := ioutil.ReadFile(filePath)
	if e != nil {
		panic(e)
	}
	slice := make([]person, 0, 1)
	s := strings.Split(string(dat), "\n")
	for _, v := range s {
		fullName := strings.Split(v, " ")
		slice = append(slice, person{fname: fullName[0], lname: fullName[1]})
	}
	for _, x := range slice {
		fmt.Printf("%s / %s\n", x.fname, x.lname)
	}
}
