package main

import "fmt"

var c, python, java bool

const Pi = 3.14

// numeric constants
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// comment out the necessary function calls to test each.
	var i int
	fmt.Println(i, c, python, java)
	// constants
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	datatype() // Call the datatype function from datatype.go
	fmt.Println("Loops Example:")
	loops() // Call the loops function from loops.go
	fmt.Println("If Example:")
	ifExample() // Call the ifExample function from if.go
	fmt.Println("Switch Example:")
	switchExample() // Call the switchExample function from switch.go
	fmt.Println("Defer Example:")
	deferExample() // Call the deferExample function from defer.go
}
