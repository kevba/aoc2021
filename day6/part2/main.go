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
	fishes := parseFish(input)
	// fishes = [8]int{0, 0, 0, 0, 0, 0, 0, 1}

	for day := 0; day < 256; day++ {
		newFishes := [9]int{}
		for timer, count := range fishes {
			// Wrap back to the 6 days timer,
			// Add new fishes to the 8 day timer
			if timer == 0 {
				newFishes[6] += count
				newFishes[8] += count
			} else {
				newFishes[timer-1] += count
			}
		}
		fishes = newFishes
	}

	answer := 0
	for _, count := range fishes {
		answer += count
	}

	return answer
}

func parseFish(input []string) (fishes [9]int) {
	fishTimers := aoc2021.AtoiSlice(strings.Split(input[0], ","))

	for _, timer := range fishTimers {
		fishes[timer] = fishes[timer] + 1
	}

	return fishes
}
