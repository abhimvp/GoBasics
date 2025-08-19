package main

import "fmt"

func loops() {
	sum := 0
	for i := 0; i < 10; i++ { //lexical scope
		sum += i
	}
	fmt.Println(sum)
	//  For continued
	sum1 := 1
	for ; sum1 < 1000; {
		sum1 += sum1
	}
	fmt.Println(sum1)
	// For is Go's "while"
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)
	// infinite loop with condition
	sum3 := 1
	for {
		if (sum3>100){
			break
		}
		sum3 += sum3
		fmt.Println("This will run forever", sum3)
	}
}