package main

import (
	"fmt"
	"math"
)

// You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a,
// initial velocity vo, and initial displacement so.
//GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement.
//The function returned by GenDisplaceFn() should take one float64 argument t,
// representing time, and return one float64 argument which is the displacement travelled after time t.

// Let us assume the following formula for displacement s as a function of time t, acceleration a,
// initial velocity vo, and initial displacement so.

// s =½ a t2 + vot + so

func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {

	fn := func(t float64) float64 {
		return ((vo * t) + (.5 * (a * math.Pow(t, 2))) + so)
	}
	return fn

}

// Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement.
// Then the program should prompt  the user to enter a value for time and the program should compute the displacement
// after the entered time.

func main() {

	var t float64
	var a float64
	var vo float64
	var so float64
	fmt.Println("Enter acceleration")
	fmt.Scan(&a)
	fmt.Println("Enter initial velocity")
	fmt.Scan(&vo)
	fmt.Println("Enter initial displacement")
	fmt.Scan(&so)
	Dist2 := GenDisplaceFn(a, vo, so)

	fmt.Println("Enter time")
	fmt.Scan(&t)
	//s := fn(t)
	fmt.Println(Dist2(t))
}
