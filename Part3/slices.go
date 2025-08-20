package main

import (
	"fmt"
	"strings"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlices(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
	// slices are like references to arrays
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	// slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	Hello := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(Hello)
	// slice defaults
	simple := []int{2, 3, 5, 7, 11, 13}

	simple = simple[1:4]
	fmt.Println(simple)

	simple = simple[:2]
	fmt.Println(simple)

	simple = simple[1:]

	fmt.Println(simple)
	// slice-len-cap

	sample := []int{2, 3, 5, 7, 11, 13}
	printSlice(sample)

	// Slice the slice to give it zero length.
	sample = sample[:0]
	printSlice(sample)

	// Extend its length.
	sample = sample[:4]
	printSlice(sample)

	// Drop its first two values.
	sample = sample[2:]
	printSlice(sample)

	// Nil Slices
	var soap []int
	fmt.Println(soap, len(soap), cap(soap))
	if soap == nil {
		fmt.Println("nil!")
	}
	// create a slice with make
	l := make([]int, 5)
	printSlices("a", l)

	m := make([]int, 0, 5)
	printSlices("b", m)

	c := m[:2]
	printSlices("c", c)

	d := c[2:5]
	printSlices("d", d)
	// slices of slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	// --- Appending to a slice ---
	var solid []int
	printSlice(solid)

	// append works on nil slices.
	solid = append(solid, 0)
	printSlice(solid)

	// The slice grows as needed.
	solid = append(solid, 1)
	printSlice(solid)

	// We can add more than one element at a time.
	solid = append(solid, 2, 3, 4)
	printSlice(solid)

}
