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
	bitLength := len(input[0])
	gammaRate := 0

	for i := 0; i < bitLength; i++ {
		if mostOnes(i, input) {
			gammaRate = (gammaRate | 1<<(bitLength-(i+1)))
		}
	}

	return gammaRate * (^gammaRate & ((1 << bitLength) - 1))
}

func mostOnes(index int, input []string) bool {
	oneCount := 0

	for _, line := range input {
		if line[index] == byte('1') {
			oneCount += 1
		}
	}

	return oneCount > len(input)/2
}
