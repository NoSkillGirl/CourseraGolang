package main

import (
	"fmt"
	"time"
)

var a int

func main() {
	a = 5
	go inc()
	go dec()
	time.Sleep(8 * time.Second)
}

func inc() {
	for i := 0; i < 10; i++ {
		a = a + 1
		fmt.Println("inc: ", a, "index: ", i)
		time.Sleep(2 * time.Second)
	}
}

func dec() {
	for i := 0; i < 10; i++ {
		a = a - 1
		fmt.Println("dec: ", a, "index: ", i)
		time.Sleep(1 * time.Second)
	}
}

// Race condition :

// A race condition is when two or more routines have access to the same resource (like variable 'a' in above code)
// and attempt to read and write to that resource without any regard to the other routines.
// This type of code can create random bugs or wrong result which is not expected from the code.
// Also, we can say that the code outcome depend on interleavings so the output will defer all the time.

// How it can occur?
//This occurs due to communication between Goroutines. When Goroutines communicate through shared variable (like
// variable 'a' in above code) the race condition occurs.

// in the above code their are 2 Goroutines. One of the Goroutines calls function which will increase the value for
// variable a by 1 and another function will decrease the value of variable a by 1. When the code is executed it prints
// different output all the time.

// Example output 1:
// dec:  5 index:  0
// inc:  6 index:  0
// dec:  4 index:  1
// inc:  5 index:  1
// dec:  4 index:  2
// dec:  3 index:  3
// inc:  4 index:  2
// dec:  3 index:  4
// dec:  2 index:  5
// inc:  3 index:  3
// dec:  2 index:  6
// dec:  1 index:  7

// Example output 1:
// dec:  5 index:  0
// inc:  6 index:  0
// dec:  4 index:  1
// dec:  3 index:  2
// inc:  4 index:  1
// dec:  3 index:  3
// inc:  4 index:  2
// dec:  3 index:  4
// dec:  2 index:  5
// inc:  3 index:  3
// dec:  2 index:  6
// dec:  1 index:  7

// Above output differs from each other as the interleavings differs i.e, we can not tell which function will be completed 1st.
// As their can be many interleavings the output will depend on it and output always differs.
