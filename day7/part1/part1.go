package main

import (
	"github.com/Vendicated/AdventOfCode2021/shared"
)

func main() {
	input, solve := shared.InitNoSplit(7, 1)
	crabPositions := shared.GetNumbers(input, ",")
	minPos, maxPos := shared.GetMinAndMax(crabPositions)
	fuelPerPos := make([]int, maxPos-minPos + 1)

	for pos := minPos; pos <= maxPos; pos++ {
		for _, crabPos := range crabPositions {
			fuelPerPos[pos - minPos] += shared.Abs(crabPos - pos)
		}
	}

	minFuel, _ := shared.GetMinAndMax(fuelPerPos)
	solve(minFuel)
}
