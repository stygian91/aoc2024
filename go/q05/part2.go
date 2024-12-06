package q05

import (
	"aoc2024/inputs"
	"fmt"
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

	for {
		if isOrdered(res, rules) {
			break
		}

		for i := 0; i < len(update)-1; i++ {
			a, b := res[i], res[i+1]
			rIdx := ruleIndex(rules, a, b)
			if rIdx == -1 {
				continue
			}

			rule := rules[rIdx]
			if rule.Low == a {
				continue
			}

			res[i], res[i+1] = res[i+1], res[i]
		}
	}

	return res
}
