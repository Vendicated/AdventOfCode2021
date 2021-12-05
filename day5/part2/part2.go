package main

import (
	"strings"

	"github.com/Vendicated/AdventOfCode2021/shared"
)

const diagramSize = 1000

func main() {
	lines, solve := shared.Init(5, 2)

	diagram := make([][]int, diagramSize)
	for i := 0; i < diagramSize; i++ {
		diagram[i] = make([]int, diagramSize)
	}

	for _, line := range lines {
		var x1, x2, y1, y2 int
		{
			parts := strings.Split(line, " -> ")
			x1, y1 = parseCoords(parts[0])
			x2, y2 = parseCoords(parts[1])
		}
		if x1 == x2 {
			left, right := getLeftAndRight(y1, y2)
			for left <= right {
				diagram[x1][left]++
				left++
			}
		} else if y1 == y2 {
			left, right := getLeftAndRight(x1, x2)
			for left <= right {
				diagram[left][y1]++
				left++
			}
		} else {
			xIncr := getIncrement(x1, x2)
			yIncr := getIncrement(y1, y2)

			for x1 != x2 && y1 != y2 {
				diagram[x1][y1]++
				x1 += xIncr
				y1 += yIncr
			}
			diagram[x1][y1]++
		}
	}

	overlapCount := 0
	for _, row := range diagram {
		for _, position := range row {
			if position > 1 {
				overlapCount++
			}
		}
	}

	solve(overlapCount)
}

func getIncrement(x1, x2 int) int {
	if x1 > x2 {
		return -1
	}
	return 1
}

func parseCoords(s string) (int, int) {
	split := strings.Split(s, ",")
	return shared.Atoi(split[0]), shared.Atoi(split[1])
}

func getLeftAndRight(x1, x2 int) (int, int) {
	if x1 > x2 {
		return x2, x1
	}
	return x1, x2
}
