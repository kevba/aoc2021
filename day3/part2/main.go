package main

import (
	"aoc2021"
	"log"
	"strconv"
)

func main() {
	defer aoc2021.Time()()
	input := aoc2021.GetInput()

	answer := solve(input)
	log.Printf("anwser: %v", answer)
}

func solve(input []string) int {
	most, least := filter(0, input)

	mostIndex := 1
	for len(most) > 1 {
		most, _ = filter(mostIndex, most)
		mostIndex += 1
	}

	leastIndex := 1
	for len(least) > 1 {
		_, least = filter(leastIndex, least)
		leastIndex += 1
	}

	o2, _ := strconv.ParseInt(most[0], 2, 64)
	co2, _ := strconv.ParseInt(least[0], 2, 64)
	return int(o2) * int(co2)
}

func filter(index int, input []string) ([]string, []string) {
	ones := []string{}
	zeroes := []string{}

	for _, line := range input {
		if line[index] == byte('1') {
			ones = append(ones, line)
		} else {
			zeroes = append(zeroes, line)
		}
	}

	if len(ones) >= len(zeroes) {
		return ones, zeroes
	}
	return zeroes, ones
}
