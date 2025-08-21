package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
} // function that takes another function as an argument

// function closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func functions() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	} // we are assigning a function to a variable - anonymous function to a variable
	fmt.Println(hypot(5, 12)) // variable is used as a function

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// closures
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
