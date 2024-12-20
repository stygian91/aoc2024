package q09

import (
	"aoc2024/inputs"
	"fmt"
	"slices"
	s "slices"
	"strings"
)

func Part1() {
	// content, err := inputs.GetInputFile("q09/demo.txt")
	content, err := inputs.GetInputFile("q09/main.txt")
	if err != nil {
		panic(err)
	}

	expanded, err := ExpandSpace(strings.TrimSpace(content))
	if err != nil {
		panic(err)
	}

	expanded = CompactPart1(expanded)
	expanded = CleanupBlocks(expanded)
	checksum := ChecksumPart1(expanded)

	fmt.Println(checksum)
}

func CompactPart1(blocks []Block) []Block {
	res := blocks
	var eIdx, dIdx int

	updateIndexes := func() {
		eIdx = s.IndexFunc(res, func(block Block) bool { return block.IsFree })
		dIdx = lastNonFreeIndex(res)
	}
	updateIndexes()

	for eIdx != -1 && dIdx != -1 && eIdx < dIdx {
		var nToMove int
		if res[eIdx].Size()-res[dIdx].Size() >= 0 {
			nToMove = res[dIdx].Size()
		} else {
			nToMove = res[eIdx].Size()
		}

		newFull := Block{Id: res[dIdx].Id, IsFree: false, Start: res[eIdx].Start, End: res[eIdx].Start + nToMove}
		res[eIdx].Start += nToMove
		res[dIdx].End -= nToMove
		eStart := res[dIdx].End + 1
		newEmpty := Block{IsFree: true, Start: eStart, End: eStart + nToMove}

		res = slices.Concat(res[:eIdx], []Block{newFull}, res[eIdx:], []Block{newEmpty})

		res = CleanupBlocks(res)
		updateIndexes()
	}

	return res
}

func ChecksumPart1(blocks []Block) int {
	sum := 0

	for _, b := range blocks {
		sum += BlockChecksum(b)
	}

	return sum
}

func lastNonFreeIndex(blocks []Block) int {
	for i := len(blocks) - 1; i >= 0; i-- {
		if !blocks[i].IsFree {
			return i
		}
	}

	return -1
}

func BlockChecksum(block Block) int {
	if block.IsFree {
		return 0
	}

	sum := 0

	for i := block.Start; i < block.End; i++ {
		sum += i * block.Id
	}

	return sum
}

func SerializeBlocks(blocks []Block) string {
	b := strings.Builder{}

	for _, block := range blocks {
		b.WriteString(block.String())
	}

	return b.String()
}
