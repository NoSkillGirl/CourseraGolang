package main

import "fmt"

func main() {
	var acc float64
	fmt.Println("Enter values for acceleration: ")
	fmt.Scanf("%f", &acc)

	var inVel float64
	fmt.Println("Enter values for initial velocity: ")
	fmt.Scanf("%f", &inVel)

	var inDis float64
	fmt.Println("Enter values for initial displacement: ")
	fmt.Scanf("%f", &inDis)

	var time float64
	fmt.Println("Enter values for initial time: ")
	fmt.Scanf("%f", &time)

	fn := GenDisplaceFn(acc, inVel, inDis)
	fmt.Println(fn(time))

}

//GenDisplaceFn function
func GenDisplaceFn(acc, inVel, inDis float64) func(float64) float64 {
	disFn := func(t float64) float64 {
		return ((0.5 * acc * t * t) + (inVel * t) + (inDis))
	}
	return disFn
}
