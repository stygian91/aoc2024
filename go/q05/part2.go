package q05

import (
	"aoc2024/inputs"
	"fmt"
	"slices"
)

func Part2() {
	// content, err := inputs.GetInputFile("q05/demo.txt")
	content, err := inputs.GetInputFile("q05/main.txt")
	if err != nil {
		panic(err)
	}

	manual := parseManual(content)
	sum := 0

	for _, update := range manual.Updates {
		if isOrdered(update, manual.Rules) {
			continue
		}

		ordered := orderUpdate(update, manual.Rules)
		sum += getMiddle(ordered)
	}

	fmt.Printf("Part 2 answer: %d\n", sum)
}

func orderUpdate(update []int, rules []Rule) []int {
	res := update[:]

	slices.SortFunc(res, func(a, b int) int {
		rIdx := ruleIndex(rules, a, b)
		if rIdx == -1 {
			return 0
		}

		rule := rules[rIdx]

		if rule.Low == a {
			return 0
		}

		return -1
	})

	return res
}
