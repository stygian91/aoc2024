package q01

import (
	"aoc2024/common/data"
	"aoc2024/inputs"
	"container/heap"
	"fmt"
	"strings"
)

func Part1() {
	dataMain, err := inputs.GetInputFile("q01/main.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(dataMain, "\n")
	lHeap := data.NewHeap()
	rHeap := data.NewHeap()

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		left, right, err := splitLine(line)
		if err != nil {
			panic(fmt.Errorf("Error on line %d: %w", i, err))
		}

		heap.Push(lHeap, left)
		heap.Push(rHeap, right)
	}

	sum := 0

	for lHeap.Len() > 0 {
		l := heap.Pop(lHeap).(int)
		r := heap.Pop(rHeap).(int)

		if l > r {
			sum += l - r
		} else {
			sum += r - l
		}
	}

	fmt.Printf("Part 1 answer: %d\n", sum)
}
