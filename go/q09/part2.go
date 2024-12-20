package q09

import (
	"aoc2024/inputs"
	"fmt"
	s "slices"
	"strings"
)

func Part2() {
	// content, err := inputs.GetInputFile("q09/demo.txt")
	content, err := inputs.GetInputFile("q09/main.txt")
	if err != nil {
		panic(err)
	}

	expanded, err := ExpandSpace(strings.TrimSpace(content))
	if err != nil {
		panic(err)
	}

	expanded = CompactPart2(expanded)
	expanded = CleanupBlocks(expanded)
	checksum := ChecksumPart1(expanded)

	fmt.Println(checksum)
}

func CompactPart2(blocks []Block) []Block {
	res := blocks

	// TODO: find max block ID and iterate over them backwards and use LastIndexFunc to find their current position in the slice
	// so we don't get OOB access errors as we're trimming the slice while iterating over it
	for i := len(res) - 1; i >= 0; i-- {
		if res[i].IsFree {
			continue
		}

		eIdx := findSpaceToTheLeft(res, i)
		if eIdx == -1 {
			continue
		}

		nToMove := res[i].Size()
		newFull := Block{
			Id:     res[i].Id,
			IsFree: false,
			Start:  res[eIdx].Start,
			End:    res[eIdx].Start + nToMove,
		}
		res[eIdx].Start += nToMove
		res[i].End -= nToMove
		eStart := res[i].End + 1
		newEmpty := Block{IsFree: true, Start: eStart, End: eStart + nToMove}

		res = s.Concat(res[:eIdx], []Block{newFull}, res[eIdx:], []Block{newEmpty})

		res = CleanupBlocks(res)
	}

	return res
}

func findSpaceToTheLeft(blocks []Block, bIdx int) int {
	bSize := blocks[bIdx].Size()

	for i, block := range blocks {
		if i >= bIdx {
			return -1
		}

		if block.IsFree && block.Size() >= bSize {
			return i
		}
	}

	return -1
}
