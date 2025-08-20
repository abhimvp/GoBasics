package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	q  = &Vertex{1, 2} // has type *Vertex
)

func structs() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4 // access and assign value to struct fields using dot notation
	fmt.Println(v.X)
	// pointer to structs
	p := &v        // Address of v is initialized to p
	p.X = 1e9      // change value of X using pointer - from 1 to 10 to the power of 9
	fmt.Println(v) // prints {1000000000 2}
	// struct literals
	fmt.Println(v1, v2, v3, q)

}
