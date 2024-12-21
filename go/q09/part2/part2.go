package part2

import (
	"aoc2024/common/slices"
	"aoc2024/inputs"
	"fmt"
	stdsl "slices"
	"strconv"
	"strings"
)

type Block struct {
	Id     int
	IsFree bool
}

func Run() {
	// content, err := inputs.GetInputFile("q09/demo.txt")
	content, err := inputs.GetInputFile("q09/main.txt")
	if err != nil {
		panic(err)
	}

	blocks, maxId := parse(content)

	for i := maxId; i >= 0; i-- {
		dLast := slices.LastIndexFunc(blocks, func(b Block) bool { return b.Id == i })
		dFirst := findDataBlockStart(blocks, dLast)
		dSize := dLast - dFirst + 1
		eFirst := findSpaceToTheLeft(blocks, dFirst, dLast)
		if eFirst == -1 {
			continue
		}

		for j := 0; j < dSize; j++ {
			blocks[eFirst+j], blocks[dFirst+j] = blocks[dFirst+j], blocks[eFirst+j]
		}
	}

	sum := 0
	for i, block := range blocks {
		if block.IsFree {
			continue
		}

		sum += i * block.Id
	}

	fmt.Printf("Part 2 answer: %d\n", sum)
}

func findSpaceToTheLeft(blocks []Block, first, last int) int {
	size := last - first + 1
	l := 0

	for i, block := range blocks {
		if i >= first {
			return -1
		}

		if block.IsFree {
			l++

			if l >= size {
				return i - size + 1
			}
		} else {
			l = 0
		}
	}

	return -1
}

func findDataBlockStart(blocks []Block, endIdx int) int {
	id := blocks[endIdx].Id
	startIdx := endIdx

	for i := endIdx; i >= 0; i-- {
		b := blocks[i]
		if b.Id != id {
			return startIdx
		}

		startIdx = i
	}

	return startIdx
}

func serialize(blocks []Block) string {
	b := strings.Builder{}

	for _, block := range blocks {
		if block.IsFree {
			b.WriteRune('.')
		} else {
			b.WriteString(strconv.Itoa(block.Id))
		}
	}

	return b.String()
}

func parse(str string) ([]Block, int) {
	id, cursor := 0, 0
	s := strings.TrimSpace(str)
	blocks := []Block{}

	for i, r := range s {
		cnt, err := strconv.Atoi(string(r))

		if i%2 == 0 {
			if err != nil {
				panic(err)
			}

			blocks = stdsl.Concat(
				blocks,
				stdsl.Repeat([]Block{{Id: id, IsFree: false}}, cnt),
			)
			id++
		} else {
			if err != nil {
				panic(err)
			}

			blocks = stdsl.Concat(
				blocks,
				stdsl.Repeat([]Block{{Id: -1, IsFree: true}}, cnt),
			)
		}

		cursor += cnt
	}

	return blocks, id - 1
}
