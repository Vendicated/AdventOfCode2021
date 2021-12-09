package main

import (
	"sort"

	"github.com/Vendicated/AdventOfCode2021/shared"
)

func main() {
	lines, solve := shared.Init(9, 2)
	nums := shared.GetAs2dArray(lines)

	rowCount, colCount := len(nums), len(nums[0])
	var basinSizes []int

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
				basinSizes = append(basinSizes, getBasinSize(&nums, rowCount, colCount, i, j))
			}
		}
	}

	sort.Ints(basinSizes)
	solve(basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3])
}

func getBasinSize(nums *[][]int, rowCount, colCount, i, j int) int {
	(*nums)[i][j] = 9
	size := 1
	fields := [][]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	for _, field := range fields {
		x, y := field[0], field[1]
		if x < 0 || y < 0 || x > rowCount-1 || y > colCount-1 {
			continue
		}
		if (*nums)[x][y] < 9 {
			size += getBasinSize(nums, rowCount, colCount, x, y)
		}
	}
	return size
}
