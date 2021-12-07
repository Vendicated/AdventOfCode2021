package main

import (
	"github.com/Vendicated/AdventOfCode2021/shared"
)

func main() {
	input, solve := shared.InitNoSplit(7, 2)
	crabPositions := shared.GetNumbers(input, ",")
	minPos, maxPos := shared.GetMinAndMax(crabPositions)
	fuelPerPos := make([]int, maxPos-minPos+1)

	for pos := minPos; pos <= maxPos; pos++ {
		for _, crabPos := range crabPositions {
			moves := shared.Abs(crabPos - pos)
			fuelPerPos[pos-minPos] += int(float32(moves) * (float32(moves + 1) / 2.0))
		}
	}

	minFuel, _ := shared.GetMinAndMax(fuelPerPos)
	solve(minFuel)
}
