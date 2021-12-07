package main

import (
	"aoc2021"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	defer aoc2021.Time()()
	input := aoc2021.GetInput()

	answer := solve(input)
	log.Printf("anwser: %v", answer)
}

func solve(input []string) int {
	pointPairs, max := parsePointsFromInput(input)
	grid := make([][]int, max[0]+1)
	for i := 0; i <= max[1]; i++ {
		grid[i] = make([]int, max[1]+1)
	}

	for _, pp := range pointPairs {
		xStep := 0
		yStep := 0
		steps := int(math.Abs(float64(pp.x1) - float64(pp.x2)))
		if steps == 0 {
			steps = int(math.Abs(float64(pp.y1) - float64(pp.y2)))
		}

		if pp.x1 > pp.x2 {
			xStep = -1
		} else if pp.x1 < pp.x2 {
			xStep = 1
		}

		if pp.y1 > pp.y2 {
			yStep = -1
		} else if pp.y1 < pp.y2 {
			yStep = 1
		}

		for i := 0; i <= steps; i++ {
			grid[pp.y1+(i*yStep)][pp.x1+(i*xStep)] += 1
		}

	}

	crossCount := 0
	for _, rows := range grid {
		for _, count := range rows {
			if count > 1 {
				crossCount++
			}
		}
	}

	return crossCount
}

func parsePointsFromInput(input []string) ([]pointPair, []int) {
	pointPairs := []pointPair{}
	max := []int{0, 0}

	for _, line := range input {
		pointStrings := strings.Split(line, " -> ")
		var x1, y1, x2, y2 int
		x1, y1, max = parsePoint(pointStrings[0], max)
		x2, y2, max = parsePoint(pointStrings[1], max)

		pp := pointPair{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		}

		pointPairs = append(pointPairs, pp)
	}

	return pointPairs, max
}

func parsePoint(pointString string, max []int) (int, int, []int) {
	points := strings.Split(pointString, ",")

	x := aoc2021.Atoi(points[0])
	y := aoc2021.Atoi(points[1])

	for i, v := range []int{x, y} {
		if v > max[i] {
			max[i] = v
		}

	}

	return x, y, max
}

type pointPair struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, val := range row {
			if val == 0 {
				fmt.Printf("%-1v", " ")
			} else {
				fmt.Printf("%-1v", val)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func getStartEnd(v1, v2 int) (int, int) {
	if v1 < v2 {
		return v1, v2
	}

	return v2, v1
}
