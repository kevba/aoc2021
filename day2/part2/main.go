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
	depth := 0
	aim := 0
	horizontalPos := 0

	moves := parseInput(input)

	for _, m := range moves {
		switch m.direction {
		case "forward":
			horizontalPos += m.distance
			depth += aim * m.distance
		case "down":
			aim += m.distance
		case "up":
			aim -= m.distance
		}
	}

	return depth * horizontalPos
}

func parseInput(input []string) []Move {
	moves := make([]Move, len(input))
	for i, line := range input {
		data := strings.Split(line, " ")

		moves[i] = Move{
			direction: data[0],
			distance:  aoc2021.Atoi(data[1]),
		}
	}

	return moves
}

type Move struct {
	direction string
	distance  int
}
