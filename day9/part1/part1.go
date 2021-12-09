package main

import (
	"github.com/Vendicated/AdventOfCode2021/shared"
)

func main() {
	lines, solve := shared.Init(9, 1)
	nums := shared.GetAs2dArray(lines)

	rowCount, colCount := len(nums), len(nums[0])
	riskLevelSum := 0

	checkHigher := func(height, i, j int) bool {
		if i < 0 || j < 0 {
			return true
		}
		if i > rowCount-1 || j > colCount-1 {
			return true
		}
		return height < nums[i][j]
	}

	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			height := nums[i][j]
			if checkHigher(height, i-1, j) && checkHigher(height, i, j-1) && checkHigher(height, i+1, j) && checkHigher(height, i, j+1) {
				riskLevelSum += 1 + height
			}
		}
	}

	solve(riskLevelSum)
}
