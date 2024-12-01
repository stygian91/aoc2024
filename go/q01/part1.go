package q01

import (
	"aoc2024/common/data"
	"container/heap"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inputs/demo.txt
var part1Demo string

//go:embed inputs/main.txt
var part1Main string

func Part1() {
	lines := strings.Split(part1Main, "\n")
	lHeap := data.New()
	rHeap := data.New()

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
