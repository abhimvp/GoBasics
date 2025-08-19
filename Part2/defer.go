package main

import "fmt"

func deferExample() {
	defer fmt.Println("world")

	fmt.Println("hello")

	// $ go run .
	// Defer Example:
	// hello
	// world

	// stacking defer
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	// $ go run .
	// Defer Example:
	// hello
	// counting
	// done
	// 9 8 7 6 5 4 3 2 1 0 - this usually should be 0 1 2 3 4 5 6 7 8 9 , since we are deferring the function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	// world
}
