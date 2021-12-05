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
	callNumbers, cards := getCardsFromInput(input)

	for _, called := range callNumbers {
		remaining := []*bingoCard{}
		for _, c := range cards {
			if !c.mark(called) {
				remaining = append(remaining, c)
			}
			if len(cards) == 1 && c.mark(called) {
				return called * c.unmarkedSum()
			}
		}
		cards = remaining
	}

	return -1
}

func newBingoCard(numbers []int) *bingoCard {
	return &bingoCard{
		numbers: numbers,
		offset:  5,
		bingo:   false,
	}
}

type bingoCard struct {
	numbers []int
	offset  int

	bingo bool
}

func (b *bingoCard) mark(calledNumber int) bool {
	for i, n := range b.numbers {
		if n == calledNumber {
			b.numbers[i] = -1

			if b.verifyRow(i / b.offset) {
				b.bingo = true
			}

			if b.verifyCol(i % b.offset) {
				b.bingo = true
			}

			break
		}
	}

	return b.bingo
}

func (b *bingoCard) unmarkedSum() int {
	sum := 0
	for _, n := range b.numbers {
		if n >= 0 {
			sum += n
		}
	}

	return sum
}

func (b *bingoCard) verifyRow(rowNum int) bool {
	start := rowNum * b.offset
	end := start + b.offset

	for _, cell := range b.numbers[start:end] {
		if cell >= 0 {
			return false
		}
	}

	return true
}

func (b *bingoCard) verifyCol(colNum int) bool {
	for i := 0; i < b.offset; i++ {
		if b.numbers[colNum+i*b.offset] >= 0 {
			return false
		}
	}

	return true
}

func getCardsFromInput(input []string) (numbers []int, cards []*bingoCard) {
	numbers = aoc2021.AtoiSlice(strings.Split(input[0], ","))

	cardNumbers := []int{}

	for i := 2; i <= len(input); i++ {

		if i == len(input) || input[i] == "" {
			cards = append(cards, newBingoCard(cardNumbers))
			cardNumbers = []int{}
			continue
		}

		line := input[i]
		parsedNumbers := []int{}
		for _, s := range strings.Split(line, " ") {
			if s != "" {
				parsedNumbers = append(parsedNumbers, aoc2021.Atoi(s))
			}
		}

		cardNumbers = append(cardNumbers, parsedNumbers...)
	}

	return numbers, cards
}
