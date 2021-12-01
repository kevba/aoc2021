package main

import (
	"aoc2021"
	"log"
)

func main() {
	defer aoc2021.Time()()
	input := aoc2021.GetInput()

	answer := solve(input)
	log.Printf("anwser: %v", answer)
}

func solve(input []string) int {
	increaseCounter := 0
	depths := aoc2021.AtoiSlice(input)

	for i, depth := range depths[1:] {
		if depth > depths[i] {
			increaseCounter++
		}
	}

	return increaseCounter
}
