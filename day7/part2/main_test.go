package main

import (
	"aoc2021"
	"testing"
)

func TestSolve(t *testing.T) {
	expectedAnswer := 168
	input := aoc2021.GetTestInput()
	answer := solve(input)

	if answer != expectedAnswer {
		t.Errorf("expected answer:  %v, got: %v", expectedAnswer, answer)
	}
}

func TestFuelUsageForPos(t *testing.T) {
	input := aoc2021.GetTestInput()
	crabPositions := parseCrabPositions(input)

	expectedAnswer := 168
	answer := fuelUsageForPos(5, crabPositions)
	if answer != expectedAnswer {
		t.Errorf("expected answer:  %v, got: %v", expectedAnswer, answer)
	}
}
