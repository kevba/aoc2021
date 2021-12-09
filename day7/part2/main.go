package main

import (
	"aoc2021"
	"log"
	"strings"
)

func main() {
	defer aoc2021.Time()()
	input := aoc2021.GetInput()

	answer := solve(input)
	log.Printf("anwser: %v", answer)
}

func solve(input []string) int {
	crabPositions := parseCrabPositions(input)

	totalFuels := make([]int, len(crabPositions))

	for i := range crabPositions {
		totalFuels[i] = fuelUsageForPos(i, crabPositions)
	}

	lowest := totalFuels[0]
	for _, f := range totalFuels {
		if f < lowest {
			lowest = f
		}
	}

	return lowest
}

func fuelUsageForPos(position int, crabs []int) int {
	fuel := 0

	for _, cp := range crabs {
		distance := aoc2021.IntAbs(position - cp)

		for i := 1; i <= distance; i++ {
			fuel += i
		}
	}

	return fuel
}

func parseCrabPositions(input []string) []int {
	crabPositions := aoc2021.AtoiSlice(strings.Split(input[0], ","))

	return crabPositions
}
