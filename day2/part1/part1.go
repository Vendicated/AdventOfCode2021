package main

import (
	"strings"

	"github.com/Vendicated/AdventOfCode2021/shared"
)

func main() {
	lines, solve := shared.Init(2, 1)

	horizontalPos, depth := 0, 0

	for _, line := range lines {
		fields := strings.Fields(line)
		i := shared.Atoi(fields[1])
		switch fields[0] {
		case "forward":
			horizontalPos += i
		case "down":
			depth += i
		case "up":
			depth -= i
		}
	}

	solve(horizontalPos * depth)
}
