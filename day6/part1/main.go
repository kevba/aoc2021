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

	for day := 0; day < 80; day++ {
		newFishes := []*Fish{}

		for _, fish := range fishes {
			newFish := fish.Spawn()
			if newFish != nil {
				newFishes = append(newFishes, newFish)
			}
		}

		fishes = append(fishes, newFishes...)
	}

	return len(fishes)
}

func parseFish(input []string) []*Fish {
	fishTimers := aoc2021.AtoiSlice(strings.Split(input[0], ","))
	fishes := make([]*Fish, len(fishTimers))

	for i, timer := range fishTimers {
		fishes[i] = newFish(timer)
	}

	return fishes
}

type Fish struct {
	timer int
}

func (f *Fish) Spawn() *Fish {
	if f.timer == 0 {
		f.timer = 6
		return newFish(8)
	}

	f.timer = f.timer - 1
	return nil
}

func newFish(timer int) *Fish {
	return &Fish{
		timer: timer,
	}
}
