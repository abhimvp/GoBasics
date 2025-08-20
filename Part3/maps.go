package main

import "fmt"

type MapsVertex struct {
	Lat, Long float64
}

var m map[string]MapsVertex // m is a map with key of string and value of MapsVertex(fields) , map is a reference type and has a ZERO value of nil for reference types
// m value is nil

var k = map[string]MapsVertex{
	"Bell Labs": MapsVertex{
		40.68433, -74.39967,
	},
	"Google": MapsVertex{
		37.42202, -122.08408,
	},
}

func maps() {
	m = make(map[string]MapsVertex) // every reference type can be initialized using make.
	m["Bell Labs"] = MapsVertex{   // key = "Bell Labs" , values = {40.68433, -74.39967}
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Map literals
	fmt.Println(k)

	// Mutating Maps
	mo := make(map[string]int)

	mo["Answer"] = 42
	fmt.Println("The value:", mo["Answer"])

	mo["Answer"] = 48
	fmt.Println("The value:", mo["Answer"])

	delete(mo, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := mo["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
