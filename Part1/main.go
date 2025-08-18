package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("My Favorite Number is", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println("Adding two numbers: ", add(3, 54))

	a, b := swap("hello", "world")
	fmt.Println("Swapped values:", a, b)

	x, y := split(17)
	fmt.Println("Split values:", x, y)
}

// output
// $ go run main.go
// Hello, 世界
//

// Adding package "math/rand"
