#!/bin/sh
set -e

dayNum="${1:?Please specify the day lol}"
day="day$dayNum"

mkdir -p "$day/part1" "$day/part2"

tee "$day/part1/part1.go" >/dev/null << EOF
package main

import (
	"github.com/Vendicated/AdventOfCode2021/shared"
)

func main() {
	lines, solve := shared.Init($dayNum, 1)
}
EOF

tee "$day/part2/part2.go" >/dev/null << EOF
package main

import (
	"github.com/Vendicated/AdventOfCode2021/shared"
)

func main() {
	lines, solve := shared.Init($dayNum, 2)
}
EOF

touch "$day/input.txt"

echo "Done! Paste the input in input.txt and you're good to go"