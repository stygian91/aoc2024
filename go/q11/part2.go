package q11

import (
	"aoc2024/inputs"
	"fmt"
	"strconv"
)

func Part2() {
	// content, err := inputs.GetInputFile("q11/demo.txt")
	content, err := inputs.GetInputFile("q11/main.txt")
	if err != nil {
		panic(err)
	}

	valueCounts := sliceToMapCount(parse(content))
	for i := 0; i < 75; i++ {
		valueCounts = processPart2(valueCounts)
	}

	cnt := 0
	for _, v := range valueCounts {
		cnt += v
	}

	fmt.Println("Part 1 answer:", cnt)
}

func sliceToMapCount(input []int) map[int]int {
	res := map[int]int{}

	for _, v := range input {
		count, exists := res[v]
		if exists {
			res[v] = count + 1
		} else {
			res[v] = 1
		}
	}

	return res
}

func processPart2(input map[int]int) map[int]int {
	res := map[int]int{}

	add := func(k, v int) {
		cnt, exists := res[k]
		if exists {
			res[k] = cnt + v
		} else {
			res[k] = v
		}
	}

	for num, count := range input {
		// rule 1
		if num == 0 {
			add(1, count)
			continue
		}

		// rule 2
		numStr := strconv.Itoa(num)
		if len(numStr)%2 == 0 {
			idx := len(numStr) / 2

			n1, err := strconv.Atoi(numStr[:idx])
			if err != nil {
				panic(err)
			}
			add(n1, count)

			n2, err := strconv.Atoi(numStr[idx:])
			if err != nil {
				panic(err)
			}
			add(n2, count)

			continue
		}

		// rule 3
		add(num*2024, count)
	}

	return res
}
