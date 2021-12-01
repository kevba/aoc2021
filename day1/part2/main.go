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
	depths := aoc2021.AtoiSlice(input)
	return deeperCounter(depths)
}

func deeperCounter(depths []int) int {
	increaseCounter := 0

	for i := range depths[0 : len(depths)-3] {
		sumCurrent := depths[i] + depths[i+1] + depths[i+2]
		sumNext := depths[i+1] + depths[i+2] + depths[i+3]

		if sumCurrent < sumNext {
			increaseCounter++
		}
	}

	return increaseCounter
}
