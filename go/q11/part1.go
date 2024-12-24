package q11

import (
	"aoc2024/inputs"
	"fmt"
	"iter"
	"slices"
	"strconv"
)

func Part1() {
	// content, err := inputs.GetInputFile("q11/demo.txt")
	content, err := inputs.GetInputFile("q11/main.txt")
	if err != nil {
		panic(err)
	}

	seq := slices.Values(parse(content))
	for i := 0; i < 25; i++ {
		seq = process(seq)
	}

	cnt := 0
	for range seq {
		cnt++
	}

	fmt.Println("Part 1 answer:", cnt)
}

func process(nums iter.Seq[int]) iter.Seq[int] {
	return func(yield func(int) bool) {
		for num := range nums {
			// rule 1
			if num == 0 {
				yield(1)
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
				yield(n1)

				n2, err := strconv.Atoi(numStr[idx:])
				if err != nil {
					panic(err)
				}
				yield(n2)

				continue
			}

			// rule 3
			yield(num * 2024)
		}
	}
}
