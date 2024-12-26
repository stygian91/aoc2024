package q12

import (
	"aoc2024/inputs"
	"fmt"
	"strings"
)

func Part1() {
	content, err := inputs.GetInputFile("q12/demo.txt")
	if err != nil {
		panic(err)
	}

	garden := parse(content)

	fmt.Println(garden)
}

type Gardens struct {
	Grid [][]rune
	// Plots

	Width, Height int
}

func parse(str string) Gardens {
	res := Gardens{}
	lines := strings.Split(strings.TrimSpace(str), "\n")
	res.Height = len(lines)

	for _, line := range lines {
		res.Width = len(line)
		row := []rune{}

		for _, r := range line {
			row = append(row, r)
		}

		res.Grid = append(res.Grid, row)
	}

	return res
}
