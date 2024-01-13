package main

import (
	"fmt"
)

func generateOutput(n int) {
	if n > 1 {
		generateOutput(n / 2)
		//  n=9 generateOutput(9/2)=> generateOutput(4)
		//  n=4 generateOutput(4/2)=> generateOutput(2)
		//  n=2 generateOutput(2/2)=> generateOutput(1) exits the if condition
	}
	// first n = 1, 2 , 4 ,9
	fmt.Print(n, " ")
}

func main() {
	input := 9
	fmt.Printf("Input: %d\tOutput: ", input)
	generateOutput(input)
}
