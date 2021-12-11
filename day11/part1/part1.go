package main

import (
	"github.com/Vendicated/AdventOfCode2021/shared"
)

var (
	rowCount int
	colCount int
)

func main() {
	lines, solve := shared.Init(11, 1)

	rowCount, colCount = len(lines), len(lines[0])

	var cavern = make([][]int, len(lines))
	result := 0
	for i, line := range lines {
		cavern[i] = shared.GetNumbers(line, "")
	}

	for cunt := 0; cunt < 100; cunt++ {
		var flashedFields [][]int

		for i := 0; i < rowCount; i++ {
			for j := 0; j < colCount; j++ {
				doField(&cavern, i, j, &result, &flashedFields)
			}
		}

		for _, field := range flashedFields {
			cavern[field[0]][field[1]] = 0
		}
	}

	solve(result)
}

func doField(cavern *[][]int, i, j int, result *int, flashedFields *[][]int) {
	(*cavern)[i][j]++
	if (*cavern)[i][j] > 9 {
		(*cavern)[i][j] = 0
		for _, flashedField := range *flashedFields {
			if flashedField[0] == i && flashedField[1] == j {
				return
			}
		}
		*flashedFields = append(*flashedFields, []int { i, j })
		*result++
		for m := -1; m < 2; m++ {
			for n := -1; n < 2; n++ {
				if m == 0 && n == 0 {
					continue
				}
				i2 := i + m
				j2 := j + n
				if i2 < 0 || j2 < 0 || i2 > rowCount - 1 || j2 > colCount - 1 {
					continue
				}
				doField(cavern, i2, j2, result, flashedFields)
			}
		}
	}
}

