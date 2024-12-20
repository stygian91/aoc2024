package q09_test

import (
	"aoc2024/inputs"
	"aoc2024/q09"
	"testing"
)

func TestExpandSpace(t *testing.T) {
	content, err := inputs.GetInputFile("q09/demo.txt")
	if err != nil {
		t.Fatal(err)
	}

	res, err := q09.ExpandSpace(content)
	if err != nil {
		t.Fatal(err)
	}

	res = q09.CleanupBlocks(res)

	exp := []q09.Block{
		{Id: 0, Start: 0, End: 2},
		{Start: 2, End: 5, IsFree: true},
		{Id: 1, Start: 5, End: 8},
		{Start: 8, End: 11, IsFree: true},
		{Id: 2, Start: 11, End: 12},
		{Start: 12, End: 15, IsFree: true},
		{Id: 3, Start: 15, End: 18},
		{Start: 18, End: 19, IsFree: true},
		{Id: 4, Start: 19, End: 21},
		{Start: 21, End: 22, IsFree: true},
		{Id: 5, Start: 22, End: 26},
		{Start: 26, End: 27, IsFree: true},
		{Id: 6, Start: 27, End: 31},
		{Start: 31, End: 32, IsFree: true},
		{Id: 7, Start: 32, End: 35},
		{Start: 35, End: 36, IsFree: true},
		{Id: 8, Start: 36, End: 40},
		{Id: 9, Start: 40, End: 42},
	}

	for i, e := range exp {
		r := res[i]
		if r != e {
			t.Errorf("Mismatch on block #%d, expected: %+v, got: %+v\n", i, e, r)
		}
	}
}

func TestCleanupBlocks(t *testing.T) {
	blocks := []q09.Block{
		{
			Id:     0,
			Start:  0,
			End:    2,
			IsFree: false,
		},
		{
			Start:  2,
			End:    2,
			IsFree: true,
		},
	}

	exp := []q09.Block{
		{
			Id:     0,
			Start:  0,
			End:    2,
			IsFree: false,
		},
	}

	blocks = q09.CleanupBlocks(blocks)

	if len(blocks) != 1 || blocks[0] != exp[0] {
		t.Errorf("Expected: %+v, got: %+v\n", exp, blocks)
	}
}
