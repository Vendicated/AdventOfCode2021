package main

import (
	"strings"

	"github.com/Vendicated/AdventOfCode2021/shared"
)

// 0 - 6
// 1 - 2  !!
// 2 - 5
// 3 - 5
// 4 - 4  !!
// 5 - 5
// 6 - 6
// 7 - 3  !!
// 8 - 7  !!
// 9 - 6

func main() {
	lines, solve := shared.Init(8, 1)

	count := 0
	for _, line := range lines {
		pipeIdx := strings.Index(line, "|")
		outputs := strings.Fields(line[pipeIdx:])
		for _, output := range outputs {
			length := len(output)
			if length == 7 || length == 3 || length == 4 || length == 2 {
				count++
			}
		}
	}

	solve(count)
}