package main

import "fmt"

var i int = 0

func assignNumber1() {
	i = 1
}

func assignNumber2() {
	i = 2
}

func main() {
	/*
		using "go functionName()" To invoke each function in a goroutine, these 2 goroutines will execute concurrently
		we have no knowledge of whether our 1st goroutine has completed or not,
		so depending on which of these 2 operations completes 1st, printed value will be either 0 or 1 or 2
		overall a race condition occurs when 2 or more threads can access shared data and they try to use it at the same time.
	*/
	go assignNumber1()
	go assignNumber2()
	fmt.Println(i)
	 /*
	 	also you can use the builtin golang "race detector" to identify the potential data races.
	 	in this case just run: >go run -race race_condition.go
	 	you will have an output similar to this one:
	 ==================
	 WARNING: DATA RACE
	 Read at 0x0000005f8138 by main goroutine:
	   main.main()
	       /go/race_condition.go:24 +0x6e

	 Previous write at 0x0000005f8138 by goroutine 7:
	   main.assignNumber1()
	       /go/race_condition.go:8 +0x3a

	 Goroutine 7 (finished) created at:
	   main.main()
	       /go/race_condition.go:22 +0x46
	 ==================
	 1
	 ==================
	 WARNING: DATA RACE
	 Write at 0x0000005f8138 by goroutine 8:
	   main.assignNumber2()
	       /go/race_condition.go:12 +0x3a

	 Previous write at 0x0000005f8138 by goroutine 7:
	   main.assignNumber1()
	       /go/race_condition.go:8 +0x3a

	 Goroutine 8 (running) created at:
	   main.main()
	       /go/race_condition.go:23 +0x5e

	 Goroutine 7 (finished) created at:
	   main.main()
	       /go/race_condition.go:22 +0x46
	 ==================
	 Found 2 data race(s)
	 exit status 66
	 */
}