package main

import "fmt"

func pointers() {
	i, j := 42, 2701

	p := &i         // point to i -> p is a pointer type
	fmt.Println(*p) // read i through the pointer - read the value of pointer using * - dereferencing operator
	*p = 21         // set i through the pointer 
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
