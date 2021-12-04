package main

import (
	"math"
	"strings"

	"github.com/Vendicated/AdventOfCode2021/shared"
)

type BingoField struct {
	fieldNum int
	isSet    bool
}

func main() {
	lines, solve := shared.Init(4, 1)

	rawMoves := strings.Split(lines[0], ",")
	moves := make([]int, len(rawMoves))
	for i, move := range rawMoves {
		moves[i] = shared.Atoi(move)
	}

	bingoSize := len(strings.Fields(lines[2]))
	leastMovesNeeded, winnerScore := math.MaxInt32, 0
	bingo := make([][]*BingoField, bingoSize)
	bingoCursor := 0

	for _, line := range lines[2:] {
		if line == "" {
			bingoCursor = 0
			checkBingo(bingo, moves, &leastMovesNeeded, &winnerScore)
			continue
		}

		currRow := make([]*BingoField, bingoSize)

		for i, num := range strings.Fields(line) {
			currRow[i] = &BingoField{fieldNum: shared.Atoi(num), isSet: false}
		}
		bingo[bingoCursor] = currRow
		bingoCursor++
	}
	// since the input is trimmed, it will have no trailing newline, so manually call check once
	checkBingo(bingo, moves, &leastMovesNeeded, &winnerScore)

	solve(winnerScore)
}

func checkBingo(bingo [][]*BingoField, moves []int, leastMovesNeeded, winnerScore *int) {
	for i, move := range moves {
	loop:
		for _, row := range bingo {
			for _, field := range row {
				if field.fieldNum == move {
					field.isSet = true
					break loop
				}
			}
		}
		if hasBingoWon(bingo) {
			if i+1 < *leastMovesNeeded {
				*leastMovesNeeded = i + 1
				*winnerScore = 0
				for _, row := range bingo {
					for _, field := range row {
						if !field.isSet {
							*winnerScore += field.fieldNum
						}
					}
				}
				*winnerScore *= move
			}
			break
		}
	}
}

func hasBingoWon(bingo [][]*BingoField) bool {
	bingoLen := len(bingo)
	// Rows
	for _, row := range bingo {
		won := true
		for _, field := range row {
			if !field.isSet {
				won = false
				break
			}
		}
		if won {
			return true
		}
	}

	// Columns
	for i := 0; i < bingoLen; i++ {
		won := true
		for j := 0; j < bingoLen; j++ {
			if !bingo[j][i].isSet {
				won = false
				break
			}
		}
		if won {
			return true
		}
	}

	return false
}
