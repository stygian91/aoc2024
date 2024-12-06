package q05

import (
	"aoc2024/inputs"
	"fmt"
)

func Part1() {
	// content, err := inputs.GetInputFile("q05/demo.txt")
	content, err := inputs.GetInputFile("q05/main.txt")
	if err != nil {
		panic(err)
	}

	manual := parseManual(content)
	orderedUpdates := [][]int{}

	for _, update := range manual.Updates {
		if !isOrdered(update, manual.Rules) {
			continue
		}

		orderedUpdates = append(orderedUpdates, update)
	}

	sum := 0

	for _, update := range orderedUpdates {
		sum += getMiddle(update)
	}

	fmt.Println(orderedUpdates)

	fmt.Printf("Part 1 answer: %d\n", sum)
}
