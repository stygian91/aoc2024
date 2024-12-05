package q05

import (
	"strconv"
	"strings"
)

type Rule struct {
	Low, High int
}

type Manual struct {
	Rules   []Rule
	Updates [][]int
}

func parseManual(str string) Manual {
	res := Manual{
		Rules:   []Rule{},
		Updates: [][]int{},
	}
	lines := strings.Split(strings.TrimSpace(str), "\n")
	inRules := true

	for _, line := range lines {
		if len(line) == 0 {
			inRules = false
			continue
		}

		if inRules {
			res.Rules = append(res.Rules, parseRule(line))
		} else {
			res.Updates = append(res.Updates, parseUpdate(line))
		}
	}

	return res
}

func parseRule(str string) Rule {
	parts := strings.Split(str, "|")

	l, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	h, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return Rule{Low: l, High: h}
}

func parseUpdate(str string) []int {
	res := []int{}
	parts := strings.Split(str, ",")

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}

		res = append(res, num)
	}

	return res
}

func isOrdered(update []int, rules []Rule) bool {
	for _, rule := range rules {
		isUpheld, _, _ := checkRule(update, rule)
		if !isUpheld {
			return false
		}
	}

	return true
}

func firstBrokenRule(update []int, rules []Rule) (int, int, int) {
	for i, rule := range rules {
		isUpheld, lowPos, highPos := checkRule(update, rule)
		if !isUpheld {
			return i, lowPos, highPos
		}
	}

	return -1, 0, 0
}

func checkRule(update []int, rule Rule) (bool, int, int) {
	foundPos := map[int]int{}

	for i, num := range update {
		if num != rule.Low && num != rule.High {
			continue
		}

		foundPos[num] = i
		lowPos, lowExists := foundPos[rule.Low]
		highPos, highExists := foundPos[rule.High]
		if !lowExists || !highExists {
			continue
		}

		if lowPos > highPos {
			return false, lowPos, highPos
		}
	}

	return true, 0, 0
}
