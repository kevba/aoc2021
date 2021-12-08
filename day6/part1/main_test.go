package main

import (
	"aoc2021"
	"testing"
)

func TestSolve(t *testing.T) {
	expectedAnswer := 5934
	input := aoc2021.GetTestInput()
	answer := solve(input)

	if answer != expectedAnswer {
		t.Errorf("expcted answer:  %v, got: %v", expectedAnswer, answer)
	}
}
