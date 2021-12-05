package main

import (
	"aoc2021"
	"testing"
)

func TestSolve(t *testing.T) {
	expectedAnswer := 1924
	input := aoc2021.GetTestInput()
	answer := solve(input)

	if answer != expectedAnswer {
		t.Errorf("expcted answer:  %v, got: %v", expectedAnswer, answer)
	}
}

func TestMark(t *testing.T) {
	for _, called := range []int{1, 2, 3, 4, 5, 6, 7, 25} {
		card := testCard()
		if card.numbers[called-1] != called {
			t.Errorf("expected %v got %v", called, card.numbers[0])
		}
		card.mark(called)
		if card.numbers[called-1] != -1 {
			t.Errorf("expected -1 got %v", card.numbers[0])
		}
	}
}

func TestUnmarkedSum(t *testing.T) {
	card := testCard()
	card.numbers = []int{
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1,
	}

	if sum := card.unmarkedSum(); sum != 20 {
		t.Errorf("expected sum to be 20, got: %v", sum)
	}

}
func TestVerifyRow(t *testing.T) {
	card := testCard()

	for i := 0; i < 5; i++ {
		if card.verifyRow(i) {
			t.Errorf("should not be valid: row number: %v card: %v", i, card.numbers)
		}
	}
	card.numbers[0] = -1
	card.numbers[1] = -1
	card.numbers[2] = -1
	card.numbers[3] = -1
	card.numbers[4] = -1

	if !card.verifyRow(0) {
		t.Errorf("should be valid: row number: %v card: %v", 0, card.numbers)
	}

	card = testCard()
	card.numbers[5] = -1
	card.numbers[6] = -1
	card.numbers[7] = -1
	card.numbers[8] = -1
	card.numbers[9] = -1

	if !card.verifyRow(1) {
		t.Errorf("should be valid: row number: %v card: %v", 0, card.numbers)
	}

}

func TestVerifyCol(t *testing.T) {
	card := testCard()

	for i := 0; i < 5; i++ {
		if card.verifyCol(i) {
			t.Errorf("should not be valid: row number: %v card: %v", i, card.numbers)
		}
	}
	card.numbers[0] = -1
	card.numbers[5] = -1
	card.numbers[10] = -1
	card.numbers[15] = -1
	card.numbers[20] = -1
	if !card.verifyCol(0) {
		t.Errorf("should be valid: col number: %v card: %v", 0, card.numbers)
	}

	card = testCard()
	card.numbers[2] = -1
	card.numbers[7] = -1
	card.numbers[12] = -1
	card.numbers[17] = -1
	card.numbers[22] = -1
	if !card.verifyCol(2) {
		t.Errorf("should be valid: col number: %v card: %v", 0, card.numbers)
	}
}

func testCard() *bingoCard {
	numbers := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	}
	return newBingoCard(numbers)
}
