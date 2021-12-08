package main

import (
	"aoc2021"
	"log"
	"math"
	"sort"
	"strings"
)

func main() {
	defer aoc2021.Time()()
	input := aoc2021.GetInput()

	answer := solve(input)
	log.Printf("anwser: %v", answer)
}

func solve(input []string) int {
	crabPos := parseCrabPositions(input)

	sorted := sort.IntSlice(crabPos)
	sorted.Sort()

	m := len(sorted) / 2
	bestPos := sorted[m]
	fuelUsage := 0

	for _, pos := range sorted {
		fuelUsage += int(math.Abs(float64(bestPos) - float64(pos)))
	}

	return fuelUsage
}

func parseCrabPositions(input []string) []int {
	crabPositions := aoc2021.AtoiSlice(strings.Split(input[0], ","))

	return crabPositions
}
