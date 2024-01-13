package main

import (
	"fmt"
)

func findMostRepeated(arr []string) string {
	frequency := make(map[string]int)

	// Count the frequency of each element in the array
	for _, element := range arr {
		frequency[element]++
	}

	// Find the element with the maximum frequency
	var mostRepeated string
	maxFrequency := 0
	for element, count := range frequency {
		if count > maxFrequency {
			maxFrequency = count
			mostRepeated = element
		}
	}

	return mostRepeated
}

func main() {
	input := []string{"apple", "pie", "apple", "apple", "apple", "red"}
	result := findMostRepeated(input)

	fmt.Printf("Input: %v\nOutput: %s\n", input, result)
}